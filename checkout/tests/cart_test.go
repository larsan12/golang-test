package tt

import (
	"context"
	checkout "route256/checkout/pkg/checkout_v1"
	loms "route256/loms/pkg/loms_v1"
	product "route256/product/pkg/product_v1"

	"github.com/golang/mock/gomock"
	"github.com/samber/lo"
)

func (s *ServiceSuite) TestGrpc() {
	s.Run("Checkout flow - success", func() {
		userId := int64(4)
		var items = []*checkout.AddToCartRequest{{User: userId, Sku: 1148162, Count: 10}, {User: userId, Sku: 1076963, Count: 5}}

		// mock external calls

		s.LomsClient.EXPECT().
			Stocks(gomock.Any(), &loms.StocksRequest{Sku: int32(items[0].Sku)}).
			Return(&loms.StocksResponse{Stocks: []*loms.StockItem{{WarehouseId: 1, Count: 40}}}, nil)

		s.LomsClient.EXPECT().
			Stocks(gomock.Any(), &loms.StocksRequest{Sku: int32(items[1].Sku)}).
			Return(&loms.StocksResponse{Stocks: []*loms.StockItem{{WarehouseId: 1, Count: 40}}}, nil)

		s.LomsClient.EXPECT().
			CreateOrder(gomock.Any(), &loms.CreateOrderRequest{User: userId, Items: []*loms.OrderItem{{Sku: items[0].Sku, Count: items[0].Count}, {Sku: items[1].Sku, Count: items[1].Count}}}).
			Return(&loms.CreateOrderResponse{OrderId: 10}, nil)

		// mock product client calls
		s.ProductClient.EXPECT().
			GetProduct(gomock.Any(), &product.GetProductRequest{Sku: items[0].Sku, Token: "test-token"}).
			Return(&product.GetProductResponse{Name: "Product 1", Price: 1000}, nil)

		s.ProductClient.EXPECT().
			GetProduct(gomock.Any(), &product.GetProductRequest{Sku: items[1].Sku, Token: "test-token"}).
			Return(&product.GetProductResponse{Name: "Product 2", Price: 2000}, nil)

		// add to card

		var err error

		_, err = s.Grpc.AddToCart(context.Background(), items[0])
		s.Require().Nil(err)
		_, err = s.Grpc.AddToCart(context.Background(), items[1])
		s.Require().Nil(err)

		// list card

		var cart *checkout.ListCartResponse
		cart, err = s.Grpc.ListCart(context.Background(), &checkout.ListCartRequest{User: userId})
		s.Require().Nil(err)

		for _, item := range cart.Items {
			obj, ok := lo.Find(items, func(i *checkout.AddToCartRequest) bool {
				return i.Sku == item.Sku
			})
			s.Require().Equal(ok, true)
			s.Require().Equal(obj.Count, item.Count)
		}

		// purchase

		_, err = s.Grpc.Puchase(context.Background(), &checkout.PuchaseRequest{User: userId})
		s.Require().Nil(err)

		cart, err = s.Grpc.ListCart(context.Background(), &checkout.ListCartRequest{User: userId})
		s.Require().Nil(err)
		s.Require().Equal(len(cart.Items), 0)
	})
}
