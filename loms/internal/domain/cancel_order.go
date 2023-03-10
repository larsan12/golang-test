package domain

import (
	"context"
)

func (m *Model) CancelOrder(ctx context.Context, orderId int64) error {
	return m.transactionManager.RunRepeteableReade(ctx, func(ctxTX context.Context) error {
		err := m.repository.CancelOrder(ctxTX, orderId)
		return err
	})
}
