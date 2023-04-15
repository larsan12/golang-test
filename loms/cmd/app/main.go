package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"route256/libs/kafka"
	"route256/libs/logger"
	"route256/libs/metrics"
	"route256/libs/workerpool"
	"route256/loms/config"
	loms_v1 "route256/loms/internal/api/v1"
	"route256/loms/internal/domain"
	"route256/loms/internal/interceptors"
	repository "route256/loms/internal/repository/postgres"
	"route256/loms/internal/repository/postgres/transactor"
	desc "route256/loms/pkg/loms_v1"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	develMode   = false
	metricsPort = "8112"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	// logger
	log := logger.New(develMode)

	// metrics init

	metrics := metrics.New("loms", metricsPort, config.ConfigData.TracesUrl, log)

	// init db pool
	pool, err := pgxpool.Connect(context.Background(), config.ConfigData.Db)
	if err != nil {
		log.Fatal("Unable to connect to database: %v\n", zap.Error(err))
	}
	defer pool.Close()

	// init db repository
	transactionManager := transactor.NewTransactionManager(pool)
	repo := repository.NewLomsRepo(transactionManager)

	// init kafka
	asyncProducer, err := kafka.NewAsyncProducer(config.ConfigData.KafkaBrokers)
	if err != nil {
		log.Fatal("kafka error", zap.Error(err))
	}

	logsSender := domain.NewOrderLogSender(log, asyncProducer, config.ConfigData.KafkaTopic, func(id string) {}, func(id string) {})

	// worker pools init
	orderCleanerWorkerPool := workerpool.NewPool[domain.Order, bool](context.Background(), 5)
	defer orderCleanerWorkerPool.Close()

	businessLogic := domain.New(log, repo, transactionManager, orderCleanerWorkerPool, logsSender)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.ConfigData.Port))
	if err != nil {
		log.Fatal("listenin error: %v", zap.Error(err))
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				metrics.MetricsServerInterceptor("grpc_server"),
				interceptors.LoggingInterceptor,
				metrics.TracingServerInterceptor(),
			),
		),
	)
	reflection.Register(server)
	desc.RegisterLomsV1Server(server, loms_v1.NewLomsV1(businessLogic))

	// run observers
	businessLogic.ObserveOldOrders(context.Background(), config.ConfigData.OrderExpirationTime)

	// metrics run http
	go metrics.Listen()

	log.Info("grpc server listening", zap.String("port", config.ConfigData.Port))
	if err = server.Serve(lis); err != nil {
		log.Fatal("serve error: %v", zap.Error(err))
	}
}
