package tt

import (
	"context"
	checkout "route256/checkout/pkg/checkout_v1"
	"route256/checkout/tests/seed"
	loms "route256/loms/pkg/loms_v1"
	product "route256/product/pkg/product_v1"

	"github.com/stretchr/testify/mock"
)

func (s *ServiceSuite) TestScenario() {
	var userId int64
	var err error

	s.Run("Success - добавить товары в корзину, оформить заказ", func() {
		/*
			Инициализируем БД и переменные
		*/
		// создаём юзера в БД
		userId, err = seed.SeedUser(s.ctx, s.Pool)
		s.Require().Nil(err)

		var items = seed.GenerateItems(userId, 5, 10)

		/*
			Положим в корзину 1-й товар
		*/

		// external mocks
		s.LomsClient.EXPECT().
			Stocks(mock.Anything, &loms.StocksRequest{Sku: int32(items[0].Sku)}).
			Return(&loms.StocksResponse{Stocks: []*loms.StockItem{{WarehouseId: 1, Count: 40}}}, nil).
			Once()

		// action
		_, err = s.Grpc.AddToCart(context.Background(), items[0])
		s.Require().Nil(err)

		/*
			Положим в корзину 2-й товар
		*/

		// external mocks
		s.LomsClient.EXPECT().
			Stocks(mock.Anything, &loms.StocksRequest{Sku: int32(items[1].Sku)}).
			Return(&loms.StocksResponse{Stocks: []*loms.StockItem{{WarehouseId: 1, Count: 40}}}, nil).
			Once()

		// action
		_, err = s.Grpc.AddToCart(context.Background(), items[1])
		s.Require().Nil(err)

		/*
			Оформим заказ
		*/

		// external mocks
		s.LomsClient.EXPECT().
			CreateOrder(mock.Anything, &loms.CreateOrderRequest{User: userId, Items: []*loms.OrderItem{{Sku: items[0].Sku, Count: items[0].Count}, {Sku: items[1].Sku, Count: items[1].Count}}}).
			Return(&loms.CreateOrderResponse{OrderId: 10}, nil).
			Once()

		s.ProductClient.EXPECT().
			GetProduct(mock.Anything, &product.GetProductRequest{Sku: items[0].Sku, Token: "testtoken"}).
			Return(&product.GetProductResponse{Name: "Product 1", Price: 1000}, nil).
			Once()

		s.ProductClient.EXPECT().
			GetProduct(mock.Anything, &product.GetProductRequest{Sku: items[1].Sku, Token: "testtoken"}).
			Return(&product.GetProductResponse{Name: "Product 2", Price: 2000}, nil).
			Once()

		// action
		_, err = s.Grpc.Puchase(context.Background(), &checkout.PuchaseRequest{User: userId})
		s.Require().Nil(err)

		/*
			Проверим что корзина пустая
		*/
		cart, err := s.Grpc.ListCart(context.Background(), &checkout.ListCartRequest{User: userId})
		s.Require().Nil(err)
		s.Require().Equal(len(cart.Items), 0)
	})
}
