package main

import (
	"context"
	"fmt"
	"log"
	"net"
	checkout_v1 "route256/checkout/internal/api/v1"
	"route256/checkout/internal/cache"
	"route256/checkout/internal/clients/grpc/loms"
	"route256/checkout/internal/clients/grpc/product"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/interceptors"
	repository "route256/checkout/internal/repository/postgres"
	"route256/checkout/internal/repository/postgres/transactor"
	desc "route256/checkout/pkg/checkout_v1"
	"route256/libs/logger"
	"route256/libs/metrics"
	"route256/libs/ratelimiter"
	"route256/libs/workerpool"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	develMode       = false
	metricsPort     = "8111"
	productCacheTtl = 30_000
)

func main() {
	// config init
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	// logger
	log := logger.New(develMode)

	// init db pool
	pool, err := pgxpool.Connect(context.Background(), config.ConfigData.Db)
	if err != nil {
		log.Fatal("Unable to connect to database: %v\n", zap.Error(err))
	}
	defer pool.Close()

	// init db repository
	transactionManager := transactor.NewTransactionManager(pool)
	repo := repository.NewCartRepo(transactionManager)

	// ratelimits
	productServiceLimiter := ratelimiter.NewLimiter(config.ConfigData.ProductServiceRateLiming, config.ConfigData.ProductServiceRateLiming)
	defer productServiceLimiter.Close()

	// metrics init

	metrics := metrics.New("checkout", metricsPort, config.ConfigData.TracesUrl, log)

	// loms client
	lomsConn, err := grpc.Dial(config.ConfigData.Services.Loms,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(metrics.MetricsClientInterceptor("loms_client")),
		grpc.WithUnaryInterceptor(metrics.TracingClientInterceptor()),
	)

	if err != nil {
		log.Fatal("failed to connect to server: %v", zap.Error(err))
	}
	defer lomsConn.Close()
	lomsClient := loms.NewClient(lomsConn)

	// product cache
	cache := cache.Create[domain.Product](productCacheTtl)

	// product client
	productConn, err := grpc.Dial(config.ConfigData.Services.Product,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(metrics.MetricsClientInterceptor("product_client")),
		grpc.WithUnaryInterceptor(metrics.TracingClientInterceptor()),
	)
	if err != nil {
		log.Fatal("failed to connect to server: %v", zap.Error(err))
	}
	defer productConn.Close()
	productClient := product.NewClient(productConn, config.ConfigData.Token, productServiceLimiter, cache)

	// pools init
	// глобальный пул для запросов к продукт сервису, вне зависимости от колличества запросов к серверу - всегда будет не более 5 паралельных запросов к продукт сервису
	getProductPool := workerpool.NewPool[uint32, domain.Product](context.Background(), config.ConfigData.GetProductPoolAmount)
	defer getProductPool.Close()

	// services init
	businessLogic := domain.New(log, lomsClient, productClient, repo, transactionManager, getProductPool)

	// server init
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
	desc.RegisterCheckoutV1Server(server, checkout_v1.NewCheckoutV1(businessLogic))
	log.Info("grpc server listening", zap.String("port", config.ConfigData.Port))
	// metrics run http
	go metrics.Listen()

	if err = server.Serve(lis); err != nil {
		log.Fatal("serve error: %v", zap.Error(err))
	}
}
