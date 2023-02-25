package loms

import (
	"context"
	"log"
	"route256/checkout/internal/domain"

	"github.com/pkg/errors"
)

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
	log.Printf("[client Stocks] sku: %d", sku)
	request := StocksRequest{SKU: sku}
	response, err := c.stocksRequest(ctx, request)

	if err != nil {
		return nil, errors.Wrap(err, "[client Stocks] stocks request error")
	}

	stocks := make([]domain.Stock, 0, len(response.Stocks))
	for _, stock := range response.Stocks {
		stocks = append(stocks, domain.Stock{
			WarehouseID: stock.WarehouseID,
			Count:       stock.Count,
		})
	}
	log.Printf("[client Stocks] %+v", stocks)

	return stocks, nil
}
