package main

import (
	"context"
	"fmt"
	"log"
	"net"
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
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}
	// init db pool
	pool, err := pgxpool.Connect(context.Background(), config.ConfigData.Db)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	// init db repository
	transactionManager := transactor.NewTransactionManager(pool)
	repo := repository.NewLomsRepo(transactionManager)

	// worker pools init
	orderCleanerWorkerPool := workerpool.NewPool[domain.Order, bool](context.Background(), 5)
	defer orderCleanerWorkerPool.Close()

	businessLogic := domain.New(repo, transactionManager, orderCleanerWorkerPool)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.ConfigData.Port))
	if err != nil {
		log.Fatalf("listenin error: %v", err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptors.LoggingInterceptor,
			),
		),
	)
	reflection.Register(server)
	desc.RegisterLomsV1Server(server, loms_v1.NewLomsV1(businessLogic))

	// run observers
	businessLogic.ObserveOldOrders(context.Background(), config.ConfigData.OrderExpirationTime)

	log.Printf("grpc server listening at %v port", config.ConfigData.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
