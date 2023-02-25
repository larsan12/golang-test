package loms

import (
	"context"
	"route256/libs/requestwrapper"
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
