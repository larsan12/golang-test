package domain

import (
	"context"
	"log"
	"route256/libs/workerpool"
)

func (m *Model) ListCart(ctx context.Context, user int64) ([]CartItem, error) {
	// TODO
	repoItems, err := m.repository.ListCart(ctx, user)
	if err != nil {
		return nil, err
	}
	items := make([]CartItem, 0, len(repoItems))

	// new context with cancel
	poolCtx, cancel := context.WithCancel(ctx)
	getProduct := func(sku uint32) Product {
		product, err := m.GetProduct(poolCtx, sku)
		if err != nil {
			log.Print("error getProduct, cancel pool", err)
			cancel()
		}
		return product
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
	products, err := m.productWorkerPool.Execute(poolCtx, tasks)

	// check error
	if poolCtx.Err() != nil {
		return nil, poolCtx.Err()
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
	log.Printf("[service ListCart] items: %+v", items)

	return items, nil
}
