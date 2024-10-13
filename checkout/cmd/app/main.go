package main

import (
	"fmt"
	"net"
	"route256/checkout/cmd/server"
	"route256/checkout/internal/config"
	"route256/libs/logger"
	"route256/libs/metrics"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	develMode       = false
	metricsPort     = "8111"
	productCacheTtl = 30_000
)

func main() {
	// logger
	log := logger.New(develMode)

	// metrics init
	metrics := metrics.New("checkout", metricsPort, config.ConfigData.TracesUrl, log)

	// loms connection
	lomsConn, err := grpc.Dial(
		config.ConfigData.Services.Loms,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(metrics.MetricsClientInterceptor("loms_client")),
		grpc.WithUnaryInterceptor(metrics.TracingClientInterceptor()),
	)
	if err != nil {
		log.Fatal("failed to connect to server: %v", zap.Error(err))
	}
	defer lomsConn.Close()

	// product connection
	productConnOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if metrics != nil {
		productConnOptions = append(
			productConnOptions,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(metrics.MetricsClientInterceptor("product_client")),
			grpc.WithUnaryInterceptor(metrics.TracingClientInterceptor()),
		)
	}
	productConn, err := grpc.Dial(config.ConfigData.Services.Product, productConnOptions...)
	if err != nil {
		log.Fatal("failed to connect to server: %v", zap.Error(err))
	}
	defer productConn.Close()

	// server init
	server, closeServer := server.Server(server.Externals{Log: log, Metrics: &metrics, LomsConn: lomsConn, ProductConn: productConn})
	defer closeServer()

	if err != nil {
		log.Fatal("config init", zapcore.Field{Type: zapcore.ErrorType, Interface: err})
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.ConfigData.Port))
	if err != nil {
		log.Fatal("listenin error: %v", zap.Error(err))
	}

	// metrics run http
	go metrics.Listen()
	log.Info("grpc server listening", zap.String("port", config.ConfigData.Port))

	if err = server.Serve(lis); err != nil {
		log.Fatal("serve error: %v", zap.Error(err))
	}
}
