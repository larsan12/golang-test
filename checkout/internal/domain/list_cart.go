package domain

import (
	"context"
	"route256/libs/workerpool"

	"go.uber.org/zap"
)

func (m *Model) ListCart(ctx context.Context, user int64) ([]CartItem, error) {
	// TODO
	repoItems, err := m.repository.ListCart(ctx, user)
	if err != nil {
		return nil, err
	}
	items := make([]CartItem, 0, len(repoItems))

	// create task
	getProduct := func(sku uint32) (Product, error) {
		return m.GetProduct(ctx, sku)
	}

	// create tasks
	tasks := make([]workerpool.Task[uint32, Product], len(repoItems))
	for i, item := range repoItems {
		tasks[i] = workerpool.Task[uint32, Product]{
			Run:    getProduct,
			InArgs: item.Sku,
		}
	}

	// execute in pool
	products, err := m.getProductPool.Execute(ctx, tasks)

	// check error
	if err != nil {
		return nil, err
	}

	// create map sku:product
	productcsMap := make(map[uint32]Product, len(products))
	for _, product := range products {
		productcsMap[product.Sku] = product
	}

	// map items
	for _, item := range repoItems {
		items = append(items, CartItem{
			Sku:   item.Sku,
			Count: item.Count,
			Name:  productcsMap[item.Sku].Name,
			Price: productcsMap[item.Sku].Price,
		})
	}
	m.log.Info("[service ListCart] items: %+v", zap.Any("items", items))

	return items, nil
}
