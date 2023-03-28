package domain_test

import (
	"context"
	"errors"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/domain/mocks"
	"route256/libs/workerpool"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestAddToCart(t *testing.T) {
	t.Parallel()

	type LomsClientMockFunc func(mc *minimock.Controller) domain.LomsClient
	type ProductServiceClientMockFunc func(mc *minimock.Controller) domain.ProductServiceClient
	type RepositoryMockFunc func(mc *minimock.Controller) domain.Repository
	type TransactionManagerMockFunc func(mc *minimock.Controller) domain.TransactionManager

	mc := minimock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx      context.Context
		cartItem domain.CartItemDiff
	}

	var (
		itemDiff1 = domain.CartItemDiff{
			User:  1,
			Sku:   1,
			Count: 10,
		}

		stock1 = domain.Stock{
			WarehouseID: 1,
			Count:       5,
		}
		stock2 = domain.Stock{
			WarehouseID: 2,
			Count:       10,
		}
		stocks    = []domain.Stock{stock1, stock2}
		someError = errors.New("some error")
	)

	tests := []struct {
		name                     string
		args                     args
		err                      error
		lomsClientMock           LomsClientMockFunc
		productServiceClientMock ProductServiceClientMockFunc
		repositoryMock           RepositoryMockFunc
		transactionManagerMock   TransactionManagerMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx:      ctx,
				cartItem: itemDiff1,
			},
			err: nil,
			lomsClientMock: func(mc *minimock.Controller) domain.LomsClient {
				mock := mocks.NewLomsClientMock(mc)
				mock.StocksMock.Expect(ctx, itemDiff1.Sku).Return(stocks, nil)
				return mock
			},
			productServiceClientMock: func(mc *minimock.Controller) domain.ProductServiceClient {
				mock := mocks.NewProductServiceClientMock(mc)
				return mock
			},
			repositoryMock: func(mc *minimock.Controller) domain.Repository {
				mock := mocks.NewRepositoryMock(mc)
				mock.GetCartItemMock.Expect(ctx, itemDiff1.User, itemDiff1.Sku).Return(itemDiff1, nil)
				mock.UpdateCartItemCountMock.Expect(ctx, itemDiff1, itemDiff1.Count*2).Return(nil)
				return mock
			},
			transactionManagerMock: func(mc *minimock.Controller) domain.TransactionManager {
				mock := mocks.NewTransactionManagerMock(mc)
				mock.RunRepeteableReadeMock.Set(func(ctx context.Context, f func(ctxTx context.Context) error) error {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "stocks error",
			args: args{
				ctx:      ctx,
				cartItem: itemDiff1,
			},
			err: errors.New("[service AddToCart] checking stocks"),
			lomsClientMock: func(mc *minimock.Controller) domain.LomsClient {
				mock := mocks.NewLomsClientMock(mc)
				mock.StocksMock.Expect(ctx, itemDiff1.Sku).Return(nil, someError)
				return mock
			},
			productServiceClientMock: func(mc *minimock.Controller) domain.ProductServiceClient {
				mock := mocks.NewProductServiceClientMock(mc)
				return mock
			},
			repositoryMock: func(mc *minimock.Controller) domain.Repository {
				mock := mocks.NewRepositoryMock(mc)
				return mock
			},
			transactionManagerMock: func(mc *minimock.Controller) domain.TransactionManager {
				mock := mocks.NewTransactionManagerMock(mc)
				return mock
			},
		},
		{
			name: "insufficient stocks",
			args: args{
				ctx:      ctx,
				cartItem: itemDiff1,
			},
			err: errors.New("insufficient stocks"),
			lomsClientMock: func(mc *minimock.Controller) domain.LomsClient {
				mock := mocks.NewLomsClientMock(mc)
				mock.StocksMock.Expect(ctx, itemDiff1.Sku).Return([]domain.Stock{stock1}, nil)
				return mock
			},
			productServiceClientMock: func(mc *minimock.Controller) domain.ProductServiceClient {
				mock := mocks.NewProductServiceClientMock(mc)
				return mock
			},
			repositoryMock: func(mc *minimock.Controller) domain.Repository {
				mock := mocks.NewRepositoryMock(mc)
				return mock
			},
			transactionManagerMock: func(mc *minimock.Controller) domain.TransactionManager {
				mock := mocks.NewTransactionManagerMock(mc)
				return mock
			},
		},
		{
			name: "repository.AddToCart errors",
			args: args{
				ctx:      ctx,
				cartItem: itemDiff1,
			},
			err: errors.New("repository.AddToCart error"),
			lomsClientMock: func(mc *minimock.Controller) domain.LomsClient {
				mock := mocks.NewLomsClientMock(mc)
				mock.StocksMock.Expect(ctx, itemDiff1.Sku).Return(stocks, nil)
				return mock
			},
			productServiceClientMock: func(mc *minimock.Controller) domain.ProductServiceClient {
				mock := mocks.NewProductServiceClientMock(mc)
				return mock
			},
			repositoryMock: func(mc *minimock.Controller) domain.Repository {
				mock := mocks.NewRepositoryMock(mc)
				mock.GetCartItemMock.Expect(ctx, itemDiff1.User, itemDiff1.Sku).Return(itemDiff1, someError)
				return mock
			},
			transactionManagerMock: func(mc *minimock.Controller) domain.TransactionManager {
				mock := mocks.NewTransactionManagerMock(mc)
				mock.RunRepeteableReadeMock.Set(func(ctx context.Context, f func(ctxTx context.Context) error) error {
					return f(ctx)
				})
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			getProductPool := workerpool.NewPool[uint32, domain.Product](ctx, 5)

			businessLogic := domain.New(tt.lomsClientMock(mc), tt.productServiceClientMock(mc), tt.repositoryMock(mc), tt.transactionManagerMock(mc), getProductPool)

			err := businessLogic.AddToCart(tt.args.ctx, tt.args.cartItem)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.Equal(t, tt.err, err)
			}
		})
	}
}
