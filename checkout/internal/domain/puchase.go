package domain

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

func (m *Model) Puchase(ctx context.Context, user int64) (int64, error) {
	var err error
	var order int64
	cartItems, err := m.ListCart(ctx, user)
	if err != nil {
		return order, errors.Wrap(err, "[service puchase] ListCart error")
	}
	log.Printf("[service puchase] cartItems: %+v", cartItems)

	orderItems := make([]OrderItem, 0, len(cartItems))

	for _, item := range cartItems {
		orderItems = append(orderItems, OrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		})
	}

	log.Printf("[service puchase] orderItems: %+v", orderItems)

	// TODO
	order, err = m.CreateOrder(ctx, user, orderItems)
	if err != nil {
		return order, errors.Wrap(err, "[service puchase] CreateOrder error")
	}

	m.repository.DeleteCart(ctx, user)
	log.Printf("[service] success puchase, orderId: %d", order)

	return order, err
}
