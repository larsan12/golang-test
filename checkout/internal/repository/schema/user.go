package schema

import "database/sql"

type User struct {
	UserId int64           `db:"user_id"`
	Name   sql.NullString  `db:"name"`
}
