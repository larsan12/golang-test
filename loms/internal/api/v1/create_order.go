package loms_v1

import (
	"context"
	"route256/loms/internal/domain"
	desc "route256/loms/pkg/loms_v1"
)

func (i *Implemetation) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	var response desc.CreateOrderResponse
	items := req.GetItems()
	var orderItems []domain.OrderItem = make([]domain.OrderItem, 0, len(items))

	for _, item := range items {
		orderItems = append(orderItems, domain.OrderItem{
			Sku:   item.Sku,
			Count: uint16(item.Count),
		})
	}

	orderId, err := i.businessLogic.CreateOrder(ctx, req.GetUser(), orderItems)
	if err != nil {
		return nil, err
	}
	response.OrderId = orderId
	return &response, nil
}
