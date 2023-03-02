package main

import (
	"fmt"
	"log"
	"net"
	checkout_v1 "route256/checkout/internal/api/v1"
	"route256/checkout/internal/clients/grpc/loms"
	"route256/checkout/internal/clients/grpc/product"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/interceptors"
	desc "route256/checkout/pkg/checkout_v1"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
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
	businessLogic := domain.New(lomsClient, productClient)

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
