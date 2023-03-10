package schema

type Stock struct {
	WarehouseId int64 `db:"warehouse_id"`
	Sku   uint32  		`db:"sku"`
	Count   uint64    `db:"count"`
}

type Order struct {
	OrderId int64     `db:"order_id"`
	UserId   uint32  		`db:"user_id"`
	Status   string   `db:"status"`
}

type StockReservation struct {
	OrderId int64 		   `db:"order_id"`
	Sku   uint32  		   `db:"sku"`
	WarehouseId int64    `db:"warehouse_id"`
	Count   uint32  		 `db:"count"`
	Status	string 			 `db:"status"`
}

type OrderItem struct {
	OrderId int64 		   `db:"order_id"`
	Sku   uint32  		   `db:"sku"`
	Count   uint32  		 `db:"count"`
}