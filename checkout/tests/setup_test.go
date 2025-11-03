package tt

import (
	"context"
	"log"
	"net"
	"os"
	"route256/checkout/cmd/server"
	"route256/checkout/internal/config"
	checkout "route256/checkout/pkg/checkout_v1"
	mock_loms_v1 "route256/checkout/tests/mocks/loms"
	mock_product_v1 "route256/checkout/tests/mocks/product"
	"route256/libs/logger"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var Psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func TestIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type ServiceSuite struct {
	suite.Suite

	// global variables
	ctx  context.Context
	Pool *pgxpool.Pool

	// server - то, что будем тестировать
	Grpc checkout.CheckoutV1Client

	// clients mocks
	LomsClient    *mock_loms_v1.LomsV1Client
	ProductClient *mock_product_v1.ProductServiceClient

	// test hooks
	beforeEach func()
}

func (s *ServiceSuite) SetupSuite() {
	s.ctx = context.Background()
	os.Setenv("IS_TEST", "true")
	if err := config.Init(); err != nil {
		panic(err)
	}

	// init db pool
	pool, err := pgxpool.Connect(context.Background(), config.ConfigData.Db)
	if err != nil {
		log.Fatal("Unable to connect to database", zap.Error(err))
	}
	s.Pool = pool
}

// SetupTest method will run before each test in the suite.
func (s *ServiceSuite) SetupSubTest() {
	// устанавливаем лок чтобы тесты выполнялись последовательно чтобы избежать конфликтов
	if _, err := s.Pool.Exec(s.ctx, "SELECT pg_advisory_lock(0)"); err != nil {
		s.T().Fatal(err)
	}

	log := logger.New(true)

	// client with mocks (mockery)
	s.LomsClient = mock_loms_v1.NewLomsV1Client(s.T())

	// product client mock (mockery)
	s.ProductClient = mock_product_v1.NewProductServiceClient(s.T())

	// init server
	server, _ := server.Server(server.Externals{Log: log, Metrics: nil, LomsClient: s.LomsClient, ProductClient: s.ProductClient, PgPool: s.Pool})

	lis := bufconn.Listen(1024 * 1024)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Info("serve error: %v", zap.Error(err))
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}

	// set grpc client
	s.Grpc = checkout.NewCheckoutV1Client(conn)

	// cleanup - close connections and release lock
	s.T().Cleanup(func() {
		conn.Close()
		lis.Close()
		//closeServer()
		if _, err := s.Pool.Exec(s.ctx, "SELECT pg_advisory_unlock(0)"); err != nil {
			s.T().Fatal(err)
		}
	})

	// зачищаем таблицы от предыдущих кейсов
	s.cleanUpTables(getTableNames()...)

	// хук, который можно использовать внутри теста через s.BeforeEach
	if s.beforeEach != nil {
		s.beforeEach()
	}
}

func (s *ServiceSuite) TearDownTest() {
	s.beforeEach = nil
}

func (s *ServiceSuite) TearDownSuite() {
	defer s.Pool.Close()
}

// хук - выполняется перед каждый тест-кейсом
func (s *ServiceSuite) BeforeEach(fn func()) {
	s.beforeEach = fn
}

func (s *ServiceSuite) cleanUpTables(tables ...string) {
	for _, tableName := range tables {
		_, err := s.Pool.Exec(s.ctx, "TRUNCATE "+tableName+" CASCADE;")
		s.Require().NoError(err)
	}
}

// for hellpers
func (s *ServiceSuite) GetCtx() context.Context {
	return s.ctx
}

func (s *ServiceSuite) GetPool() *pgxpool.Pool {
	return s.Pool
}

// список таблиц для зачистки перед/после каждого тест кейса
func getTableNames() []string {
	res := []string{
		"users",
		"cart_items",
	}

	return res
}
