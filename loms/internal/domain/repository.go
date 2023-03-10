package domain

import "context"

type Repository interface {
	MakeOrder(ctx context.Context, user int64, items []OrderItem) (int64, error)
	ListOrder(ctx context.Context, orderId int64) (Order, error)
	PaidOrder(ctx context.Context, orderId int64) error
	CancelOrder(ctx context.Context, orderId int64) error
	Stock(ctx context.Context, sku uint32) ([]StockItem, error)
}

