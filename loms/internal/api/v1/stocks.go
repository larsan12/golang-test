package loms_v1

import (
	"context"
	desc "route256/loms/pkg/loms_v1"
)

func (i *Implemetation) Stocks(ctx context.Context, req *desc.StocksRequest) (*desc.StocksResponse, error) {
	var response desc.StocksResponse

	items, err := i.businessLogic.Stock(ctx, uint32(req.Sku))
	if err != nil {
		return nil, err
	}
	var resItems []*desc.StockItem = make([]*desc.StockItem, 0, len(items))

	for _, item := range items {
		resItems = append(resItems, &desc.StockItem{
			WarehouseId: item.WarehouseID,
			Count:       item.Count,
		})
	}

	response.Stocks = resItems
	return &response, nil
}
