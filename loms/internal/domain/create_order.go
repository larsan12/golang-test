package domain

import (
	"context"
	"errors"
)

func (m *Model) CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error) {
	var orderId int64
	var err error
	m.transactionManager.RunRepeteableReade(ctx, func(ctxTX context.Context) error {
		orderId, err = m.repository.CreateOrder(ctx, user)
		if err != nil {
			return err
		}
		for _, item := range items {
			stocks, err := m.repository.GetStocks(ctxTX, item.Sku)
			if err != nil {
				return err;
			}
			restCount := item.Count
			for _, stock := range stocks {
				var reservation StockReservation = StockReservation{
					OrderId: orderId,
					WarehouseId: stock.WarehouseID,
					Sku: item.Sku,
				}
				reservation.WarehouseId = stock.WarehouseID
				if stock.Count >= uint64(restCount) {
					reservation.Count = uint32(restCount)
					restCount = 0
				} else {
					restCount = restCount - uint16(stock.Count)
					reservation.Count = uint32(stock.Count)
				}
				err = m.repository.CreateStockReservation(ctxTX, reservation)
				if err != nil {
					return err
				}
				stock.Count -= uint64(reservation.Count)
				err = m.repository.UpdateStock(ctxTX, item.Sku, stock.WarehouseID, -int64(reservation.Count))
				if err != nil {
					return err
				}
				if restCount == 0 {
					break;
				}
			}
			if restCount != 0 {
				return errors.New("not enough stocks")
			}
		}
		return nil
	})
	return orderId, err
}
