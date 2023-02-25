package domain

import (
	"context"
)

func (m *Model) ListOrder(ctx context.Context, orderId int64) (Order, error) {
	order := [2]int32{1076963, 1148162}
	items := make([]OrderItem, 0, len(order))
	for _, sku := range order {
		items = append(items, OrderItem{
			Sku:   uint32(sku),
			Count: 1,
		})
	}

	return Order{
		Status: "new",
		User:   1,
		Items:  items,
	}, nil
}
