package tt

import (
	"context"
	"net"
	"os"
	"route256/checkout/cmd/server"
	"route256/checkout/internal/config"
	checkout "route256/checkout/pkg/checkout_v1"
	mock_loms_v1 "route256/checkout/tests/mocks"
	"route256/libs/logger"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func Init(m *testing.M) {
	// before all tests
	close := initGrpcServer()
	code := m.Run()
	// after all tests
	close()
	os.Exit(code)
}

var (
	Grpc checkout.CheckoutV1Client
	Psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	Pool *pgxpool.Pool
	lis  *bufconn.Listener

	// loms client
	lomsController *gomock.Controller
	LomsClient     *mock_loms_v1.MockLomsV1Client
)

func initGrpcServer() func() {
	log := logger.New(true)

	// init db pool
	pool, err := pgxpool.Connect(context.Background(), config.ConfigData.Db)
	if err != nil {
		log.Fatal("Unable to connect to database", zap.Error(err))
	}
	Pool = pool

	// loms client
	lomsController = gomock.NewController(&testing.T{})
	LomsClient = mock_loms_v1.NewMockLomsV1Client(lomsController)

	// product connection
	productConn, err := grpc.NewClient(config.ConfigData.Services.Product, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to server: %v", zap.Error(err))
	}

	// server init
	s, closeServer := server.Server(server.Externals{Log: log, Metrics: nil, LomsClient: LomsClient, ProductConn: productConn, PgPool: pool})

	if err != nil {
		log.Fatal("config init", zapcore.Field{Type: zapcore.ErrorType, Interface: err})
	}

	bufSize := 1024 * 1024
	lis = bufconn.Listen(bufSize)
	go func() {
		if err = s.Serve(lis); err != nil {
			log.Info("serve error: %v", zap.Error(err))
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
		pool.Close()
	}
}
