package domain

import (
	"context"
)

func (m *Model) ListOrder(ctx context.Context, orderId int64) (Order, error) {
	var order Order
	var err error
	err = m.transactionManager.RunRepeteableReade(ctx, func(ctxTX context.Context) error {
		order, err = m.repository.GetOrder(ctxTX, orderId)
		if err != nil {
			return err
		}
		orderItems, err := m.repository.GetOrderItems(ctxTX, orderId)
		if err != nil {
			return err
		}
		order.Items = orderItems
		return nil
	})
	return order, err
}
