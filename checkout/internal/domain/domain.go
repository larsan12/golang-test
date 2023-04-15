package domain

import (
	"context"
	"route256/libs/workerpool"

	"go.uber.org/zap"
)

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i LomsClient -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i ProductServiceClient -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i Repository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i TransactionManager -o ./mocks/ -s "_minimock.go"

type LomsClient interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
	CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error)
}

type ProductServiceClient interface {
	Product(ctx context.Context, sku uint32) (Product, error)
}

type TransactionManager interface {
	RunRepeteableReade(ctx context.Context, f func(ctxTX context.Context) error) error
}

type Model struct {
	log                  *zap.Logger
	lomsClient           LomsClient
	productServiceClient ProductServiceClient
	repository           Repository
	transactionManager   TransactionManager

	getProductPool workerpool.WorkerPool[uint32, Product]
}

func New(
	log *zap.Logger,
	lomsClient LomsClient,
	productServiceClient ProductServiceClient,
	repository Repository,
	transactionManager TransactionManager,
	getProductPool workerpool.WorkerPool[uint32, Product],
) *Model {
	return &Model{
		log,
		lomsClient,
		productServiceClient,
		repository,
		transactionManager,
		getProductPool,
	}
}

type Stock struct {
	WarehouseID int64
	Count       uint64
}

type OrderItem struct {
	Sku   uint32
	Count uint16
}

type Product struct {
	Name  string
	Price uint32
	Sku   uint32
}

type CartItemDiff struct {
	User  int64
	Sku   uint32
	Count uint16
}

type CartItem struct {
	Sku   uint32
	Count uint16
	Name  string
	Price uint32
}
