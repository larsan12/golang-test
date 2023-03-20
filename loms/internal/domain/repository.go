package domain

import (
	"context"
	"time"
)

type Repository interface {
	CreateOrder(ctx context.Context, user int64) (int64, error)
	CreateStockReservation(ctx context.Context, stockReserv StockReservation) error
	UpdateStock(ctx context.Context, sku uint32, warehouseId int64, diff int64) error
	GetStocks(ctx context.Context, sku uint32) ([]StockItem, error)
	GetStockReservations(ctx context.Context, orderId int64) ([]StockReservation, error)
	OrderSetStatus(ctx context.Context, orderId int64, status string) error
	DeleteStockReservation(ctx context.Context, stockReserv StockReservation) error
	ReservSetStatuses(ctx context.Context, orderId int64, status string) error
	GetOrder(ctx context.Context, orderId int64) (Order, error)
	GetOrderItems(ctx context.Context, orderId int64) ([]OrderItem, error)
	GetOldOrders(ctx context.Context, dateFrom time.Time) ([]Order, error)
}
