package product

import (
	"context"
	"fmt"
	"route256/checkout/internal/domain"
	"route256/libs/cache"
	"route256/libs/ratelimiter"
	productServiceAPI "route256/product/pkg/product_v1"
	"time"

	"google.golang.org/grpc"
)

var _ domain.ProductServiceClient = (*client)(nil)

// ProductServiceClientInterface defines the interface for product service client
type ProductServiceClientInterface interface {
	GetProduct(ctx context.Context, in *productServiceAPI.GetProductRequest, opts ...grpc.CallOption) (*productServiceAPI.GetProductResponse, error)
	ListSkus(ctx context.Context, in *productServiceAPI.ListSkusRequest, opts ...grpc.CallOption) (*productServiceAPI.ListSkusResponse, error)
}

// Client ...
type client struct {
	productClient         ProductServiceClientInterface
	token                 string
	productServiceLimiter ratelimiter.Limiter
	cache                 cache.Cache[domain.Product]
}

// NewClient creates a new product client
func NewClient(clientInterface ProductServiceClientInterface, token string, productServiceLimiter ratelimiter.Limiter, cache cache.Cache[domain.Product]) *client {
	return &client{
		productClient:         clientInterface,
		token:                 token,
		productServiceLimiter: productServiceLimiter,
		cache:                 cache,
	}
}

func (c *client) Product(ctx context.Context, sku uint32) (domain.Product, error) {
	if val, ok := c.cache.Get(fmt.Sprint(sku)); ok {
		return val, nil
	}
	// ratelimiting
	c.productServiceLimiter.Wait(ctx)

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
	product.Sku = sku
	c.cache.Set(fmt.Sprint(sku), product)
	return product, nil
}
