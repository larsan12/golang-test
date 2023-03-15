package main

import (
	"context"
	"fmt"
	"log"
	"net"
	checkout_v1 "route256/checkout/internal/api/v1"
	"route256/checkout/internal/clients/grpc/loms"
	"route256/checkout/internal/clients/grpc/product"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/interceptors"
	repository "route256/checkout/internal/repository/postgres"
	"route256/checkout/internal/repository/postgres/transactor"
	desc "route256/checkout/pkg/checkout_v1"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	// config init
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
	repo := repository.NewCartRepo(transactionManager)

	// loms client
	lomsConn, err := grpc.Dial(config.ConfigData.Services.Loms, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer lomsConn.Close()
	lomsClient := loms.NewClient(lomsConn)

	// product client
	productConn, err := grpc.Dial(config.ConfigData.Services.Product, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer productConn.Close()
	productClient := product.NewClient(productConn, config.ConfigData.Token)

	// services init
	businessLogic := domain.New(lomsClient, productClient, repo, transactionManager)

	// server init
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
	desc.RegisterCheckoutV1Server(server, checkout_v1.NewCheckoutV1(businessLogic))
	log.Printf("grpc server listening at %v port", config.ConfigData.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
