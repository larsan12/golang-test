package domain

import (
	"context"
)

func (m *Model) OrderPayed(ctx context.Context, orderId int64) error {
	return m.transactionManager.RunRepeteableReade(ctx, func(ctxTX context.Context) error {
		err := m.repository.OrderSetStatus(ctxTX, orderId, OrderStatusPaid)
		if err != nil {
			return err
		}
		err = m.repository.ReservSetStatuses(ctxTX, orderId, ReserveStatusPaid)
		if err != nil {
			return err
		}
		return nil
	})
}
