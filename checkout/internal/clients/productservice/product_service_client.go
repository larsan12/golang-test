package productservice

import (
	"context"
	"route256/libs/requestwrapper"
)

type Client struct {
	url               string
	token             string
	getProductRequest func(context.Context, GetProductRequest) (GetProductResponse, error)
}

func New(url string, token string) *Client {
	return &Client{
		url:               url,
		token:             token,
		getProductRequest: requestwrapper.PostRequester[GetProductRequest, GetProductResponse](url + "/get_product"),
	}
}
