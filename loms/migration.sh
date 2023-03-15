goose -dir ./migrations postgres "postgres://user:password@localhost:6432/loms?sslmode=disable" status

goose -dir ./migrations postgres "postgres://user:password@localhost:6432/loms?sslmode=disable" up
