package main

import (
	"fmt"
	"log"
	"net"
	"route256/loms/config"
	loms_v1 "route256/loms/internal/api/v1"
	"route256/loms/internal/domain"
	"route256/loms/internal/interceptors"
	desc "route256/loms/pkg/loms_v1"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	businessLogic := domain.New()

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
	log.Printf("grpc server listening at %v port", config.ConfigData.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
