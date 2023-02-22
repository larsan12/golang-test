package loms

import (
	"context"
	"log"
	"route256/checkout/internal/domain"

	"github.com/pkg/errors"
)

type CreateOrderRequest struct {
	User  int64             `json:"user"`
	Items []CreateOrderItem `json:"items"`
}

type CreateOrderItem struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

func (c *Client) CreateOrder(ctx context.Context, user int64, items []domain.OrderItem) (int64, error) {
	log.Printf("[client CreateOrder] user: %d ; items: %+v", user, items)
	orderItems := make([]CreateOrderItem, 0, len(items))
	for _, item := range items {
		orderItems = append(orderItems, CreateOrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		})
	}
	request := CreateOrderRequest{User: user, Items: orderItems}
	response, err := c.createOrderRequest(ctx, request)

	if err != nil {
		return 0, errors.Wrap(err, "[client CreateOrder] stocks request error")
	}
	log.Printf("[client CreateOrder] orderId: %d", response.OrderID)

	return response.OrderID, nil
}
