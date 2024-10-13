package tests

import (
	"context"
	checkout "route256/checkout/pkg/checkout_v1"
	"route256/checkout/tests"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	tests.Init(m)
}

func TestGrpc(t *testing.T) {
	t.Run("Add to cart - success", func(t *testing.T) {
		var items = []*checkout.AddToCartRequest{{User: 4, Sku: 1148162, Count: 10}, {User: 4, Sku: 1076963, Count: 5}}
		var err error
		_, err = tests.Grpc.AddToCart(context.Background(), items[0])
		require.NoError(t, err)
		_, err = tests.Grpc.AddToCart(context.Background(), items[1])
		require.NoError(t, err)

		var cart *checkout.ListCartResponse
		cart, err = tests.Grpc.ListCart(context.Background(), &checkout.ListCartRequest{User: 4})
		require.NoError(t, err)

		for _, item := range cart.Items {
			obj, ok := lo.Find(items, func(i *checkout.AddToCartRequest) bool {
				return i.Sku == item.Sku
			})
			require.Equal(t, ok, true)
			require.Equal(t, obj.Count, item.Count)
		}
	})
}
