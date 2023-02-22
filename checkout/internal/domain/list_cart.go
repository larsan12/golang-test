package domain

import (
	"context"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

func (m *Model) ListCart(ctx context.Context, user int64) ([]CartItem, error) {
	// TODO
	cart := [2]int32{1076963, 1148162}
	items := make([]CartItem, 0, len(cart))

	for _, sku := range cart {
		info, err := m.GetProduct(ctx, uint32(sku))
		if err != nil {
			return items, errors.Wrap(err, "[service ListCart] getProduct error, sku: "+strconv.Itoa(int(sku)))
		}
		items = append(items, CartItem{
			Sku:   uint32(sku),
			Count: 1,
			Name:  info.Name,
			Price: info.Price,
		})
	}
	log.Printf("[service ListCart] items: %+v", items)

	return items, nil
}
