package domain

import (
	"context"
	"route256/libs/workerpool"
)

type Model struct {
	repository             Repository
	transactionManager     TransactionManager
	orderCleanerWorkerPool workerpool.WorkerPool[Order, bool]
	logsSender             LogsSender
}

type LogsSender interface {
	SendOrderStatusAsync(order Order)
}

func New(repository Repository, transactionManager TransactionManager, orderCleanerWorkerPool workerpool.WorkerPool[Order, bool], logsSender LogsSender) *Model {
	return &Model{
		repository,
		transactionManager,
		orderCleanerWorkerPool,
		logsSender,
	}
}

type TransactionManager interface {
	RunRepeteableReade(ctx context.Context, f func(ctxTX context.Context) error) error
}

type OrderItem struct {
	Sku   uint32
	Count uint16
}

type Order struct {
	OrderId int64 `json:"id"`
	Status  string
	User    int64       `json:"-"`
	Items   []OrderItem `json:"-"`
}

type StockItem struct {
	WarehouseID int64
	Count       uint64
}

type StockReservation struct {
	OrderId     int64
	Sku         uint32
	WarehouseId int64
	Count       uint32
	Status      string
}

const (
	OrderStatusNew      = "new"
	OrderStatusPaid     = "paid"
	OrderStatusCanceled = "cancelled"
	ReserveStatusPaid   = "paid"
)
