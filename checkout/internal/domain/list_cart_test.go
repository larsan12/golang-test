package domain_test

import (
	"context"
	"errors"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/domain/mocks"
	"route256/libs/logger"
	"route256/libs/workerpool"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestListCart(t *testing.T) {
	t.Parallel()

	type LomsClientMockFunc func(mc *minimock.Controller) domain.LomsClient
	type ProductServiceClientMockFunc func(mc *minimock.Controller) domain.ProductServiceClient
	type RepositoryMockFunc func(mc *minimock.Controller) domain.Repository
	type TransactionManagerMockFunc func(mc *minimock.Controller) domain.TransactionManager

	mc := minimock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx  context.Context
		user int64
	}

	var (
		userId    int64 = 1
		itemDiff1       = domain.CartItemDiff{
			Sku:   1,
			Count: 10,
		}
		itemDiff2 = domain.CartItemDiff{
			Sku:   2,
			Count: 20,
		}
		itemDiffList = []domain.CartItemDiff{itemDiff1, itemDiff2}
		product1     = domain.Product{
			Sku:   1,
			Name:  "p1",
			Price: 15,
		}
		product2 = domain.Product{
			Sku:   2,
			Name:  "p2",
			Price: 25,
		}

		item1 = domain.CartItem{
			Sku:   itemDiff1.Sku,
			Count: itemDiff1.Count,
			Name:  product1.Name,
			Price: product1.Price,
		}
		item2 = domain.CartItem{
			Sku:   itemDiff2.Sku,
			Count: itemDiff2.Count,
			Name:  product2.Name,
			Price: product2.Price,
		}
		itemList  = []domain.CartItem{item1, item2}
		someError = errors.New("some error")
	)

	tests := []struct {
		name                     string
		args                     args
		want                     []domain.CartItem
		err                      error
		lomsClientMock           LomsClientMockFunc
		productServiceClientMock ProductServiceClientMockFunc
		repositoryMock           RepositoryMockFunc
		transactionManagerMock   TransactionManagerMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx:  ctx,
				user: userId,
			},
			want: itemList,
			err:  nil,
			lomsClientMock: func(mc *minimock.Controller) domain.LomsClient {
				mock := mocks.NewLomsClientMock(mc)
				return mock
			},
			productServiceClientMock: func(mc *minimock.Controller) domain.ProductServiceClient {
				mock := mocks.NewProductServiceClientMock(mc)
				mock.ProductMock.When(ctx, product1.Sku).Then(product1, nil)
				mock.ProductMock.When(ctx, product2.Sku).Then(product2, nil)
				return mock
			},
			repositoryMock: func(mc *minimock.Controller) domain.Repository {
				mock := mocks.NewRepositoryMock(mc)
				mock.ListCartMock.Expect(ctx, userId).Return(itemDiffList, nil)
				return mock
			},
			transactionManagerMock: func(mc *minimock.Controller) domain.TransactionManager {
				mock := mocks.NewTransactionManagerMock(mc)
				return mock
			},
		},
		{
			name: "repository error",
			args: args{
				ctx:  ctx,
				user: userId,
			},
			want: nil,
			err:  someError,
			lomsClientMock: func(mc *minimock.Controller) domain.LomsClient {
				mock := mocks.NewLomsClientMock(mc)
				return mock
			},
			productServiceClientMock: func(mc *minimock.Controller) domain.ProductServiceClient {
				mock := mocks.NewProductServiceClientMock(mc)
				return mock
			},
			repositoryMock: func(mc *minimock.Controller) domain.Repository {
				mock := mocks.NewRepositoryMock(mc)
				mock.ListCartMock.Expect(ctx, userId).Return(nil, someError)
				return mock
			},
			transactionManagerMock: func(mc *minimock.Controller) domain.TransactionManager {
				mock := mocks.NewTransactionManagerMock(mc)
				return mock
			},
		},
		{
			name: "productService - one request failed with error - returns error",
			args: args{
				ctx:  ctx,
				user: userId,
			},
			want: nil,
			err:  errors.New("[service GetProduct] getp product"),
			lomsClientMock: func(mc *minimock.Controller) domain.LomsClient {
				mock := mocks.NewLomsClientMock(mc)
				return mock
			},
			productServiceClientMock: func(mc *minimock.Controller) domain.ProductServiceClient {
				mock := mocks.NewProductServiceClientMock(mc)
				mock.ProductMock.When(ctx, product1.Sku).Then(product1, nil)
				mock.ProductMock.When(ctx, product2.Sku).Then(domain.Product{}, someError)
				return mock
			},
			repositoryMock: func(mc *minimock.Controller) domain.Repository {
				mock := mocks.NewRepositoryMock(mc)
				mock.ListCartMock.Expect(ctx, userId).Return(itemDiffList, nil)
				return mock
			},
			transactionManagerMock: func(mc *minimock.Controller) domain.TransactionManager {
				mock := mocks.NewTransactionManagerMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			log := logger.New(true)
			getProductPool := workerpool.NewPool[uint32, domain.Product](ctx, 5)

			businessLogic := domain.New(log, tt.lomsClientMock(mc), tt.productServiceClientMock(mc), tt.repositoryMock(mc), tt.transactionManagerMock(mc), getProductPool)

			res, err := businessLogic.ListCart(tt.args.ctx, tt.args.user)
			require.Equal(t, tt.want, res)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.Equal(t, tt.err, err)
			}
		})
	}
}
