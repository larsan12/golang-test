package domain

import (
	"context"
)

func (m *Model) DeleteFromCart(ctx context.Context, cartItem CartItemDiff) error {
	err := m.transactionManager.RunRepeteableReade(ctx, func(ctxTX context.Context) error {
		err := m.repository.DeleteFromCart(ctx, cartItem)
		return err
	})
	return err
}
