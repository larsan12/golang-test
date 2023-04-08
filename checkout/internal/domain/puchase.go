package domain

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (m *Model) Puchase(ctx context.Context, user int64) (int64, error) {
	var err error
	var order int64
	cartItems, err := m.ListCart(ctx, user)
	if err != nil {
		return order, errors.Wrap(err, "[service puchase] ListCart error")
	}
	m.log.Info("[service puchase]", zap.Any("cartItems", cartItems))

	orderItems := make([]OrderItem, 0, len(cartItems))

	for _, item := range cartItems {
		orderItems = append(orderItems, OrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		})
	}

	m.log.Info("[service puchase]", zap.Any("orderItems", orderItems))

	// TODO
	order, err = m.CreateOrder(ctx, user, orderItems)
	if err != nil {
		return order, errors.Wrap(err, "[service puchase] CreateOrder error")
	}

	m.repository.DeleteCart(ctx, user)
	m.log.Info("[service] success puchase, orderId: %d", zap.Int64("orderId", order))

	return order, err
}
