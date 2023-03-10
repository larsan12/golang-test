package domain

import (
	"context"
)

func (m *Model) Stock(ctx context.Context, sku uint32) ([]StockItem, error) {
	return m.repository.Stock(ctx, sku)
}
