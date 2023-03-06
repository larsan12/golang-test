package loms

import (
	"context"
	"route256/checkout/internal/domain"
	lomsServiceAPI "route256/loms/pkg/loms_v1"
	"time"

	"google.golang.org/grpc"
)

var _ domain.LomsClient = (*client)(nil)

// Client ...
type client struct {
	lomsClient lomsServiceAPI.LomsV1Client
}

// NewClient ...
func NewClient(cc *grpc.ClientConn) *client {
	return &client{
		lomsClient: lomsServiceAPI.NewLomsV1Client(cc),
	}
}

func (c *client) Stocks(ctx context.Context, sku uint32) ([]domain.Stock, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := c.lomsClient.Stocks(ctx, &lomsServiceAPI.StocksRequest{
		Sku: int32(sku),
	})

	if err != nil {
		return nil, err
	}
	resStocks := res.GetStocks()
	var stocks []domain.Stock = make([]domain.Stock, 0, len(resStocks))

	for _, item := range resStocks {
		stocks = append(stocks, domain.Stock{
			WarehouseID: item.WarehouseId,
			Count:       item.Count,
		})
	}

	return stocks, nil
}

func (c *client) CreateOrder(ctx context.Context, user int64, items []domain.OrderItem) (int64, error) {
	var orderItems []*lomsServiceAPI.OrderItem = make([]*lomsServiceAPI.OrderItem, 0, len(items))

	for _, item := range items {
		orderItems = append(orderItems, &lomsServiceAPI.OrderItem{
			Sku:   item.Sku,
			Count: uint32(item.Count),
		})
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := c.lomsClient.CreateOrder(ctx, &lomsServiceAPI.CreateOrderRequest{
		User:  user,
		Items: orderItems,
	})

	if err != nil {
		return 0, err
	}

	return res.OrderId, nil
}
