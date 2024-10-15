package tt

import (
	"context"
	checkout "route256/checkout/pkg/checkout_v1"
	tt "route256/checkout/tests"
	loms "route256/loms/pkg/loms_v1"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	tt.Init(m)
}

func TestGrpc(t *testing.T) {
	t.Run("Checkout flow - success", tt.Run(func(t *testing.T) {
		userId := int64(4)
		var items = []*checkout.AddToCartRequest{{User: userId, Sku: 1148162, Count: 10}, {User: userId, Sku: 1076963, Count: 5}}

		// mock external calls

		tt.LomsClient.EXPECT().
			Stocks(gomock.Any(), &loms.StocksRequest{Sku: int32(items[0].Sku)}).
			Return(&loms.StocksResponse{Stocks: []*loms.StockItem{{WarehouseId: 1, Count: 40}}}, nil)

		tt.LomsClient.EXPECT().
			Stocks(gomock.Any(), &loms.StocksRequest{Sku: int32(items[1].Sku)}).
			Return(&loms.StocksResponse{Stocks: []*loms.StockItem{{WarehouseId: 1, Count: 40}}}, nil)

		tt.LomsClient.EXPECT().
			CreateOrder(gomock.Any(), &loms.CreateOrderRequest{User: userId, Items: []*loms.OrderItem{{Sku: items[0].Sku, Count: items[0].Count}, {Sku: items[1].Sku, Count: items[1].Count}}}).
			Return(&loms.CreateOrderResponse{OrderId: 10}, nil)

		// add to card

		var err error

		_, err = tt.Grpc.AddToCart(context.Background(), items[0])
		require.NoError(t, err)
		_, err = tt.Grpc.AddToCart(context.Background(), items[1])
		require.NoError(t, err)

		// list card

		var cart *checkout.ListCartResponse
		cart, err = tt.Grpc.ListCart(context.Background(), &checkout.ListCartRequest{User: userId})
		require.NoError(t, err)

		for _, item := range cart.Items {
			obj, ok := lo.Find(items, func(i *checkout.AddToCartRequest) bool {
				return i.Sku == item.Sku
			})
			require.Equal(t, ok, true)
			require.Equal(t, obj.Count, item.Count)
		}

		// purchase

		_, err = tt.Grpc.Puchase(context.Background(), &checkout.PuchaseRequest{User: userId})
		require.NoError(t, err)

		cart, err = tt.Grpc.ListCart(context.Background(), &checkout.ListCartRequest{User: userId})
		require.NoError(t, err)
		require.Equal(t, len(cart.Items), 0)

	}))
}
