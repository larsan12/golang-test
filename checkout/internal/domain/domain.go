package domain

import "context"

type LomsService interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
	CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error)
}

type ProductService interface {
	Product(ctx context.Context, sku uint32) (Product, error)
}

type Model struct {
	lomsService    LomsService
	productService ProductService
}

func New(lomsService LomsService, productService ProductService) *Model {
	return &Model{
		lomsService:    lomsService,
		productService: productService,
	}
}
