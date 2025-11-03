package seed

import (
	"context"
	"math/rand/v2"
	checkout "route256/checkout/pkg/checkout_v1"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

var Psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// SeedUser creates a user in the database and returns the user ID
func SeedUser(ctx context.Context, pool *pgxpool.Pool) (int64, error) {
	query, args, err := Psql.
		Insert("users").
		Columns("name").
		Values("Test User").
		Suffix("RETURNING user_id").
		ToSql()
	if err != nil {
		return 0, err
	}

	var userID int64
	err = pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

type CartItem struct {
	Sku   uint32
	Count uint16
}

// SeedCart ensures user exists and inserts cart items for the user (upsert on conflict)
func SeedCart(ctx context.Context, pool *pgxpool.Pool, userID int64, items []CartItem) error {
	// ensure user exists (allow passing existing id)
	qUser, aUser, err := Psql.
		Insert("users").
		Columns("user_id", "name").
		Values(userID, "Seed User").
		Suffix("ON CONFLICT (user_id) DO NOTHING").
		ToSql()
	if err != nil {
		return err
	}
	if _, err := pool.Exec(ctx, qUser, aUser...); err != nil {
		return err
	}

	if len(items) == 0 {
		return nil
	}

	ins := Psql.Insert("cart_items").Columns("user_id", "sku", "count")
	for _, it := range items {
		ins = ins.Values(userID, it.Sku, it.Count)
	}
	qCart, aCart, err := ins.Suffix("ON CONFLICT (user_id, sku) DO UPDATE SET count = EXCLUDED.count").ToSql()
	if err != nil {
		return err
	}
	if _, err := pool.Exec(ctx, qCart, aCart...); err != nil {
		return err
	}

	return nil
}

// генерирует айтемы в корзине
func GenerateItems(userId int64, counts ...uint32) []*checkout.AddToCartRequest {
	result := make([]*checkout.AddToCartRequest, len(counts))
	for i, v := range counts {
		result[i] = &checkout.AddToCartRequest{User: userId, Sku: rand.Uint32(), Count: v}
	}

	return result
}
