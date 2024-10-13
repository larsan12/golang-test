package tests

import (
	"context"
	"net"
	"os"
	"route256/checkout/cmd/server"
	"route256/checkout/internal/config"
	checkout "route256/checkout/pkg/checkout_v1"
	"route256/libs/logger"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func Init(m *testing.M) {
	// before test
	close := initGrpcServer()
	code := m.Run()
	// after test
	close()
	os.Exit(code)
}

var Grpc checkout.CheckoutV1Client

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func initGrpcServer() func() {
	log := logger.New(true)

	// loms connection
	lomsConn, err := grpc.NewClient(
		config.ConfigData.Services.Loms,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("failed to connect to server: %v", zap.Error(err))
	}

	// product connection
	productConn, err := grpc.NewClient(config.ConfigData.Services.Product, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to server: %v", zap.Error(err))
	}

	// server init
	s, closeServer := server.Server(server.Externals{Log: log, Metrics: nil, LomsConn: lomsConn, ProductConn: productConn})

	if err != nil {
		log.Fatal("config init", zapcore.Field{Type: zapcore.ErrorType, Interface: err})
	}

	log.Info("grpc server listening", zap.String("port", config.ConfigData.Port))

	lis = bufconn.Listen(bufSize)
	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatal("serve error: %v", zap.Error(err))
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}

	Grpc = checkout.NewCheckoutV1Client(conn)

	return func() {
		conn.Close()
		lis.Close()
		closeServer()
	}
}
