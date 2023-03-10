package domain

import (
	"context"
)

func (m *Model) CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error) {
	var orderId int64
	var err error
	m.transactionManager.RunRepeteableReade(ctx, func(ctxTX context.Context) error {
		orderId, err = m.repository.MakeOrder(ctxTX, user, items)
		return err
	})
	return orderId, err
}
