package domain

import (
	"context"
)

func (m *Model) CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error) {
	return 10, nil
}
