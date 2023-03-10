package domain

import (
	"context"
	"log"
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
