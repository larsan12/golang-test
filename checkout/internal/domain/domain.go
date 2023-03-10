package domain

import "context"

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
	lomsClient           LomsClient
	productServiceClient ProductServiceClient
	repository Repository
	transactionManager TransactionManager
}

func New(lomsClient LomsClient, productServiceClient ProductServiceClient, repository Repository, transactionManager TransactionManager) *Model {
	return &Model{
		lomsClient:           lomsClient,
		productServiceClient: productServiceClient,
		repository: repository,
		transactionManager: transactionManager,
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
