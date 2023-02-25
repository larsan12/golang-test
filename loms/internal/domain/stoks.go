package domain

import (
	"context"
)

func (m *Model) Stock(ctx context.Context, sku uint32) ([]StockItem, error) {
	return []StockItem{
		{
			WarehouseID: 123,
			Count:       5,
		},
	}, nil
}
