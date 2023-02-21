package loms

import (
	"context"
	"log"
	"route256/checkout/internal/domain"
	"route256/libs/requestwrapper"

	"github.com/pkg/errors"
)

type Client struct {
	url                string
	stocksRequest      func(context.Context, StocksRequest) (StocksResponse, error)
	createOrderRequest func(context.Context, CreateOrderRequest) (CreateOrderResponse, error)
}

func New(url string) *Client {
	return &Client{
		url:                url,
		stocksRequest:      requestwrapper.PostRequester[StocksRequest, StocksResponse](url + "/stocks"),
		createOrderRequest: requestwrapper.PostRequester[CreateOrderRequest, CreateOrderResponse](url + "/createOrder"),
	}
}

// Stock Request

type StocksRequest struct {
	SKU uint32 `json:"sku"`
}

type StocksItem struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type StocksResponse struct {
	Stocks []StocksItem `json:"stocks"`
}

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]domain.Stock, error) {
	log.Printf("[client] Stocks, sku: %d", sku)
	request := StocksRequest{SKU: sku}
	response, err := c.stocksRequest(ctx, request)

	if err != nil {
		return nil, errors.Wrap(err, "stocks request error")
	}

	stocks := make([]domain.Stock, 0, len(response.Stocks))
	for _, stock := range response.Stocks {
		stocks = append(stocks, domain.Stock{
			WarehouseID: stock.WarehouseID,
			Count:       stock.Count,
		})
	}
	log.Printf("[client] Stocks: %+v", stocks)

	return stocks, nil
}

// CreateOrder Request

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
	log.Printf("[client] CreateOrder, user: %d ; items: %+v", user, items)
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
		return 0, errors.Wrap(err, "[client] stocks request error")
	}
	log.Printf("[client] CreateOrder, orderId: %d", response.OrderID)

	return response.OrderID, nil
}
