package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m *Model) DeleteFromCart(ctx context.Context, cartItem CartItemDiff) error {
	err := m.transactionManager.RunRepeteableReade(ctx, func(ctxTX context.Context) error {
		item, err := m.repository.GetCartItem(ctxTX, cartItem.User, cartItem.Sku)
		if err != nil {
			return err
		}
		if item.Count < cartItem.Count {
			return errors.New("bad request")
		}
		return m.repository.UpdateCartItemCount(ctxTX, cartItem, item.Count - cartItem.Count)
	})
	return err
}
