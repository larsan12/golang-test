package domain

import (
	"context"
)

type BusinessLogic interface {
	AddToCart(ctx context.Context, cartItem CartItemDiff) error
	Puchase(ctx context.Context, user int64) (int64, error)
	DeleteFromCart(ctx context.Context, cartItem CartItemDiff) error
	ListCart(ctx context.Context, user int64) ([]CartItem, error)
}
