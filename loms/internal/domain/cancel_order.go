package domain

import (
	"context"
)

func (m *Model) CancelOrder(ctx context.Context, orderId int64) error {
	return m.transactionManager.RunRepeteableReade(ctx, func(ctxTX context.Context) error {
		err := m.repository.OrderSetStatus(ctxTX, orderId, OrderStatusCanceled)
		if err != nil {
			return err
		}
		reservations, err := m.repository.GetStockReservations(ctxTX, orderId)
		if err != nil {
			return err
		}

		for _, reserv := range reservations {
			err = m.repository.UpdateStock(ctxTX, reserv.Sku, reserv.WarehouseId, int64(reserv.Count))
			if err != nil {
				return err
			}
			err = m.repository.DeleteStockReservation(ctxTX, reserv)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
