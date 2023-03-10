package domain

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

var (
	ErrInsufficientStocks = errors.New("insufficient stocks")
	ErrAddToCartRepository = errors.New("repository.AddToCart error")
)

func (m *Model) AddToCart(ctx context.Context, cartItem CartItemDiff) error {
	stocks, err := m.lomsClient.Stocks(ctx, cartItem.Sku)
	if err != nil {
		return errors.WithMessage(err, "[service AddToCart] checking stocks")
	}

	counter := int64(cartItem.Count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {

			// in transaction
			err := m.transactionManager.RunRepeteableReade(ctx, func(ctxTX context.Context) error {
				err := m.repository.AddToCart(ctxTX, cartItem)
				return err
			})

			if err != nil {
				log.Printf("failed repository.AddToCart: %v", err)
				return ErrAddToCartRepository
			}
			return nil
		}
	}

	return ErrInsufficientStocks
}
