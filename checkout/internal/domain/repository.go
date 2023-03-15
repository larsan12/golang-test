package domain

import (
	"context"
)

type Repository interface {
	DeleteCart(ctx context.Context, user int64) error
	DeleteFromCart(ctx context.Context, cartItem CartItemDiff) error
	ListCart(ctx context.Context, user int64) ([]CartItemDiff, error)
	GetCartItem(ctx context.Context, user int64, sku uint32) (CartItemDiff, error)
	CreateCartItem(ctx context.Context, cartItem CartItemDiff) error
	UpdateCartItemCount(ctx context.Context, cartItem CartItemDiff, newCount uint16) error
}

