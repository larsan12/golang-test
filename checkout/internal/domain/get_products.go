package domain

import (
	"context"

	"github.com/pkg/errors"
)

type Product struct {
	Name  string
	Price uint32
}

func (m *Model) GetProduct(ctx context.Context, sku uint32) (Product, error) {
	product, err := m.productService.Product(ctx, sku)
	if err != nil {
		return product, errors.WithMessage(err, "getp product")
	}

	return product, nil
}
