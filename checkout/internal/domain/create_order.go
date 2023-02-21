package domain

import (
	"context"

	"github.com/pkg/errors"
)

type OrderItem struct {
	Sku   uint32
	Count uint16
}

func (m *Model) CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error) {
	orderId, err := m.lomsService.CreateOrder(ctx, user, items)
	if err != nil {
		return orderId, errors.WithMessage(err, "order creation request error")
	}

	return orderId, nil
}
