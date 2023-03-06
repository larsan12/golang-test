package checkout_v1

import (
	"context"
	desc "route256/checkout/pkg/checkout_v1"
)

func (i *Implemetation) ListCart(ctx context.Context, req *desc.ListCartRequest) (*desc.ListCartResponse, error) {
	var response desc.ListCartResponse

	items, err := i.businessLogic.ListCart(ctx, req.User)
	if err != nil {
		return nil, err
	}

	var resItems []*desc.CartItem = make([]*desc.CartItem, 0, len(items))

	var totalPrice uint32
	for _, item := range items {
		totalPrice += item.Price
		resItems = append(resItems, &desc.CartItem{
			Sku:   item.Sku,
			Count: uint32(item.Count),
			Price: item.Price,
			Name:  item.Name,
		})
	}

	response.TotalPrice = totalPrice
	response.Items = resItems

	return &response, nil
}
