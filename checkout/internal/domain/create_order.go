package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m *Model) CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error) {
	orderId, err := m.lomsClient.CreateOrder(ctx, user, items)
	if err != nil {
		return orderId, errors.WithMessage(err, "[service CreateOrder] order creation request error")
	}

	return orderId, nil
}
