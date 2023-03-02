package loms_v1

import (
	"context"
	desc "route256/loms/pkg/loms_v1"
)

func (i *Implemetation) ListOrder(ctx context.Context, req *desc.ListOrderRequest) (*desc.ListOrderResponse, error) {
	var response desc.ListOrderResponse

	order, err := i.businessLogic.ListOrder(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	var resItems []*desc.OrderItem = make([]*desc.OrderItem, 0, len(order.Items))

	for _, item := range order.Items {
		resItems = append(resItems, &desc.OrderItem{
			Sku:   item.Sku,
			Count: uint32(item.Count),
		})
	}

	response.Status = order.Status
	response.User = order.User
	response.Items = resItems
	return &response, nil
}
