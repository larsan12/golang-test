package domain

type Model struct {
}

func New() *Model {
	return &Model{}
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
