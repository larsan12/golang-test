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
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Order struct {
	Status string      `json:"status"`
	User   int64       `json:"user"`
	Items  []OrderItem `json:"items"`
}

type StockItem struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}
