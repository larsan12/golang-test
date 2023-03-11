package repository

import (
	"context"
	"errors"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/repository/postgres/transactor"
	"route256/checkout/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type CartRepo struct {
	transactor.QueryEngineProvider
}

func NewCartRepo(provider transactor.QueryEngineProvider) *CartRepo {
	return &CartRepo{
		QueryEngineProvider: provider,
	}
}

var (
	cartItemsColumns = []string{"user_id", "sku", "count"}
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

const (
	cartItemsTable = "cart_items"
)

func (r *CartRepo) GetCartItem(ctx context.Context, user int64, sku uint32) (domain.CartItemDiff, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var schemaItem schema.CartItem
	var item domain.CartItemDiff
	rawQuery, args, err := psql.Select(cartItemsColumns...).
		From(cartItemsTable).
		Where(sq.Eq{"user_id": user, "sku": sku}).
		ToSql()
	
	if err != nil {
		return item, err
	}
	error := pgxscan.Get(ctx, db, &schemaItem, rawQuery, args...)
	item.User = schemaItem.UserId
	item.Count = schemaItem.Count
	item.Sku = schemaItem.Sku
	return item, error
}

func (r *CartRepo) CreateCartItem(ctx context.Context, cartItem domain.CartItemDiff) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Insert(cartItemsTable).
		Columns(cartItemsColumns...).
		Values(cartItem.User, cartItem.Sku, cartItem.Count).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := db.Exec(ctx, rawQuery, args...); err != nil {
		return err
	}
	return nil
}

func (r *CartRepo) UpdateCartItemCount(ctx context.Context, cartItem domain.CartItemDiff, newCount uint16) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Update(cartItemsTable).
		Where(sq.Eq{"user_id": cartItem.User, "sku": cartItem.Sku}).
		Set("count", newCount).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := db.Exec(ctx, rawQuery, args...); err != nil {
		return err
	}
	return nil
}


func (r *CartRepo) DeleteFromCart(ctx context.Context, cartItem domain.CartItemDiff) error {
	item, err := r.GetCartItem(ctx, cartItem.User, cartItem.Sku)
	if err != nil {
		return err
	}
	if item.Count < cartItem.Count {
		return errors.New("bad request")
	}
	return r.UpdateCartItemCount(ctx, cartItem, item.Count - cartItem.Count)
}

func (r *CartRepo) DeleteCart(ctx context.Context, user int64) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Delete(cartItemsTable).
		Where(sq.Eq{"user_id": user}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = db.Exec(ctx, rawQuery, args...)
	return err
}

func (r *CartRepo) ListCart(ctx context.Context, user int64) ([]domain.CartItemDiff, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Select(cartItemsColumns...).
		From(cartItemsTable).
		Where(sq.Eq{"user_id": user}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var items []schema.CartItem
	if err := pgxscan.Select(ctx, db, &items, rawQuery, args...); err != nil {
		return nil, err
	}

	return bindSchemaItemsToModelsItems(items), nil
}

func bindSchemaItemsToModelsItems(items []schema.CartItem) []domain.CartItemDiff {
	result := make([]domain.CartItemDiff, len(items))
	for i := range items {
		result[i].User = items[i].UserId
		result[i].Count = items[i].Count
		result[i].Sku = items[i].Sku
	}
	return result
}

