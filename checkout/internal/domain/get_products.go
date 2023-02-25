package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m *Model) GetProduct(ctx context.Context, sku uint32) (Product, error) {
	product, err := m.productServiceClient.Product(ctx, sku)
	if err != nil {
		return product, errors.WithMessage(err, "[service GetProduct] getp product")
	}

	return product, nil
}
