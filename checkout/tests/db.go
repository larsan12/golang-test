package tt

import "context"

const (
	cartItemsTable = "cart_items"
)

func CleanCartItemTable() error {
	rawQuery, args, err := Psql.Delete(cartItemsTable).
		ToSql()
	if err != nil {
		return err
	}
	_, err = Pool.Exec(context.Background(), rawQuery, args...)
	return err
}
