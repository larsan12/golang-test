package server

import (
	"context"
	"log"
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
	"route256/libs/metrics"
	"route256/libs/ratelimiter"
	"route256/libs/workerpool"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	productCacheTtl = 30_000
)

type Externals struct {
	Log         *zap.Logger
	Metrics     *metrics.InterceptorsMetricsFactory
	LomsConn    *grpc.ClientConn
	ProductConn *grpc.ClientConn
}

func Server(externals Externals) (*grpc.Server, func()) {
	// init db pool
	pool, err := pgxpool.Connect(context.Background(), config.ConfigData.Db)
	if err != nil {
		log.Fatal("Unable to connect to database", zap.Error(err))
	}

	// init db repository
	transactionManager := transactor.NewTransactionManager(pool)
	repo := repository.NewCartRepo(transactionManager)

	// ratelimits
	productServiceLimiter := ratelimiter.NewLimiter(config.ConfigData.ProductServiceRateLiming, config.ConfigData.ProductServiceRateLiming)
	defer productServiceLimiter.Close()

	metrics := externals.Metrics

	// product cache
	cache := cache.Create[domain.Product](productCacheTtl)

	// loms client
	lomsClient := loms.NewClient(externals.LomsConn)

	// product client
	productClient := product.NewClient(externals.ProductConn, config.ConfigData.Token, productServiceLimiter, cache)

	// pools init
	// глобальный пул для запросов к продукт сервису, вне зависимости от колличества запросов к серверу - всегда будет не более 5 паралельных запросов к продукт сервису
	getProductPool := workerpool.NewPool[uint32, domain.Product](context.Background(), config.ConfigData.GetProductPoolAmount)

	// services init
	businessLogic := domain.New(externals.Log, lomsClient, productClient, repo, transactionManager, getProductPool)

	// server init
	serverOptions := []grpc.UnaryServerInterceptor{interceptors.LoggingInterceptor}
	if metrics != nil {
		serverOptions = append(serverOptions, (*metrics).MetricsServerInterceptor("grpc_server"), (*metrics).TracingServerInterceptor())
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(serverOptions...),
		),
	)
	reflection.Register(server)
	desc.RegisterCheckoutV1Server(server, checkout_v1.NewCheckoutV1(businessLogic))

	return server, func() {
		getProductPool.Close()
		pool.Close()
	}
}
