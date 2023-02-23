package handlers

import (
	"context"
	"route256/loms/internal/domain"
)


type BusinessLogic interface {
	CancelOrder(ctx context.Context, orderId int64) error
	CreateOrder(ctx context.Context, user int64, items []domain.OrderItem) (int64, error)
	ListOrder(ctx context.Context, orderId int64) (domain.Order, error)
	OrderPayed(ctx context.Context, orderId int64) error
	Stock(ctx context.Context, sku uint32) ([]domain.StockItem, error)
}
