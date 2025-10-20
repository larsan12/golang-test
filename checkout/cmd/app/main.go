package main

import (
	"context"
	"fmt"
	"net"
	"route256/checkout/cmd/server"
	"route256/checkout/internal/config"
	"route256/libs/logger"
	"route256/libs/metrics"

	lomsServiceAPI "route256/loms/pkg/loms_v1"
	productServiceAPI "route256/product/pkg/product_v1"

	"github.com/jackc/pgx/v4/pgxpool"
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

	// init db pool
	pool, err := pgxpool.Connect(context.Background(), config.ConfigData.Db)
	if err != nil {
		log.Fatal("Unable to connect to database", zap.Error(err))
	}
	defer pool.Close()

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
	productClient := productServiceAPI.NewProductServiceClient(productConn)

	// server init
	server, closeServer := server.Server(server.Externals{Log: log, Metrics: &metrics, LomsClient: lomsServiceAPI.NewLomsV1Client(lomsConn), ProductClient: productClient, PgPool: pool})
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
