package domain

import "context"

type Model struct {
	repository Repository
	transactionManager TransactionManager
}

func New(repository Repository, transactionManager TransactionManager) *Model {
	return &Model{
		repository,
		transactionManager,
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
	Status string
	User   int64
	Items  []OrderItem
}

type StockItem struct {
	WarehouseID int64
	Count       uint64
}

type StockReservation struct {
	OrderId 		int64
	Sku   			uint32
	WarehouseId int64
	Count   		uint32
	Status			string
}

const (
	OrderStatusNew = "new"
	OrderStatusPaid = "paid"
	OrderStatusCanceled = "cancelled"
	ReserveStatusPaid = "paid"
)
