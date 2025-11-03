package tt

import (
	"context"
	checkout "route256/checkout/pkg/checkout_v1"
	"route256/checkout/tests/seed"
	product "route256/product/pkg/product_v1"

	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
)

func (s *ServiceSuite) TestApi() {
	var userId int64
	var err error

	s.BeforeEach(func() {
		// создаём юзера в БД
		userId, err = seed.SeedUser(s.ctx, s.Pool)
		s.Require().Nil(err)
	})

	s.Run("Success - получить корзину", func() {
		/*
			Инициализируем БД и константы
		*/
		var items = seed.GenerateItems(userId, 5, 10)

		err = seed.SeedCart(s.ctx, s.Pool, userId, []seed.CartItem{
			{Sku: uint32(items[0].Sku), Count: uint16(items[0].Count)},
			{Sku: uint32(items[1].Sku), Count: uint16(items[1].Count)},
		})
		s.Require().Nil(err)

		/*
			мокаем внешние вызовы
		*/
		s.ProductClient.EXPECT().
			GetProduct(mock.Anything, &product.GetProductRequest{Sku: items[0].Sku, Token: "testtoken"}).
			Return(&product.GetProductResponse{Name: "Product 1", Price: 1000}, nil).
			Once()

		s.ProductClient.EXPECT().
			GetProduct(mock.Anything, &product.GetProductRequest{Sku: items[1].Sku, Token: "testtoken"}).
			Return(&product.GetProductResponse{Name: "Product 2", Price: 2000}, nil).
			Once()

		/*
			вызов API
		*/
		cart, err := s.Grpc.ListCart(context.Background(), &checkout.ListCartRequest{User: userId})
		s.Require().Nil(err)

		/*
			checks
		*/
		for _, item := range cart.Items {
			obj, ok := lo.Find(items, func(i *checkout.AddToCartRequest) bool {
				return i.Sku == item.Sku
			})
			s.Require().Equal(ok, true)
			s.Require().Equal(obj.Count, item.Count)
		}
	})

	s.Run("Success - получить корзину с фильтром по productName", func() {
		// ...
	})

	s.Run("Fail - не верный пользователь", func() {
		// ...
	})
}
