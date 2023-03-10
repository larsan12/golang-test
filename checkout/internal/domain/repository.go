package domain

import "context"

type Repository interface {
	AddToCart(ctx context.Context, cartItem CartItemDiff) error
	DeleteCart(ctx context.Context, user int64) error
	DeleteFromCart(ctx context.Context, cartItem CartItemDiff) error
	ListCart(ctx context.Context, user int64) ([]CartItemDiff, error)
}

