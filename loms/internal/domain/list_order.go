package domain

import (
	"context"
)

func (m *Model) ListOrder(ctx context.Context, orderId int64) (Order, error) {
	return m.repository.ListOrder(ctx, orderId)
}
