package product

import (
	"context"
	"route256/checkout/internal/domain"
	productServiceAPI "route256/checkout/pkg/product_v1"
	"time"

	"google.golang.org/grpc"
)

var _ domain.ProductServiceClient = (*client)(nil)

// Client ...
type client struct {
	productClient productServiceAPI.ProductServiceClient
	token         string
}

// NewClient ...
func NewClient(cc *grpc.ClientConn, token string) *client {
	return &client{
		productClient: productServiceAPI.NewProductServiceClient(cc),
		token:         token,
	}
}

func (c *client) Product(ctx context.Context, sku uint32) (domain.Product, error) {
	var product domain.Product
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := c.productClient.GetProduct(ctx, &productServiceAPI.GetProductRequest{
		Sku:   sku,
		Token: c.token,
	})

	if err != nil {
		return product, err
	}
	product.Name = res.GetName()
	product.Price = res.GetPrice()

	return product, nil
}
