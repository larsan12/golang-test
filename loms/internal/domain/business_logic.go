package domain

import (
	"context"
)

type BusinessLogic interface {
	CancelOrder(ctx context.Context, orderId int64) error
	CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error)
	ListOrder(ctx context.Context, orderId int64) (Order, error)
	OrderPayed(ctx context.Context, orderId int64) error
	Stock(ctx context.Context, sku uint32) ([]StockItem, error)
	ObserveOldOrders(ctx context.Context)
}
