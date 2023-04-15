package domain

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrInsufficientStocks  = errors.New("insufficient stocks")
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
				item, err := m.repository.GetCartItem(ctxTX, cartItem.User, cartItem.Sku)
				if err != nil {
					if err.Error() == "scanning one: no rows in result set" {
						return m.repository.CreateCartItem(ctxTX, cartItem)
					}
					return err
				} else {
					return m.repository.UpdateCartItemCount(ctxTX, cartItem, item.Count+cartItem.Count)
				}
			})

			if err != nil {
				m.log.Info("failed repository.AddToCart: %v", zap.Error(err))
				return ErrAddToCartRepository
			}
			return nil
		}
	}

	return ErrInsufficientStocks
}
