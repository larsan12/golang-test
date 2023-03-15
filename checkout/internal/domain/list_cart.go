package domain

import (
	"context"
	"log"
	"route256/libs/workerpool"
	"strconv"

	"github.com/pkg/errors"
)

func (m *Model) ListCart(ctx context.Context, user int64) ([]CartItem, error) {
	// TODO
	repoItems, err := m.repository.ListCart(ctx, user)
	if err != nil {
		return nil, err
	}
	items := make([]CartItem, 0, len(repoItems))

	getProduct := func(sku uint32) (Product, error) {
		var product Product
		var err error
		product, err = m.GetProduct(ctx, sku)
		if err != nil {

		}
		return product, nil
	}

	tasks = make([]workerpool.Task[uint32, Product], len(items))

	products := m.productWorkerPool.Execute()

	for _, item := range repoItems {
		info, err := m.GetProduct(ctx, item.Sku)
		if err != nil {
			return items, errors.Wrap(err, "[service ListCart] getProduct error, sku: "+strconv.Itoa(int(item.Sku)))
		}
		items = append(items, CartItem{
			Sku:   item.Sku,
			Count: item.Count,
			Name:  info.Name,
			Price: info.Price,
		})
	}
	log.Printf("[service ListCart] items: %+v", items)

	return items, nil
}
