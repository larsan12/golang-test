package domain

import "context"

type LomsClient interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
	CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error)
}

type ProductServiceClient interface {
	Product(ctx context.Context, sku uint32) (Product, error)
}

type Model struct {
	lomsClient           LomsClient
	productServiceClient ProductServiceClient
}

func New(lomsClient LomsClient, productServiceClient ProductServiceClient) *Model {
	return &Model{
		lomsClient:           lomsClient,
		productServiceClient: productServiceClient,
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

type CartItem struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}
