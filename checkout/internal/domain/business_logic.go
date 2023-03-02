package domain

import (
	"context"
)

type BusinessLogic interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error
	Puchase(ctx context.Context, user int64) (int64, error)
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
	ListCart(ctx context.Context, user int64) ([]CartItem, error)
}
