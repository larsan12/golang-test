package productservice

import (
	"context"
	"route256/checkout/internal/domain"
	"route256/libs/requestwrapper"

	"github.com/pkg/errors"
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

type GetProductRequest struct {
	Token string `json:"token"`
	SKU   uint32 `json:"sku"`
}

type GetProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (c *Client) Product(ctx context.Context, sku uint32) (domain.Product, error) {
	request := GetProductRequest{SKU: sku, Token: c.token}
	var result domain.Product
	response, err := c.getProductRequest(ctx, request)
	if err != nil {
		return result, errors.Wrap(err, "[client] Product request error")
	}
	result.Name = response.Name
	result.Price = response.Price

	return result, nil
}
