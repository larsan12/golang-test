package domain

import (
	"context"
)

func (m *Model) OrderPayed(ctx context.Context, orderId int64) error {
	return m.repository.PaidOrder(ctx, orderId)
}
