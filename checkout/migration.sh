goose -dir ./migrations postgres "postgres://user:password@localhost:6433/checkout?sslmode=disable" status

goose -dir ./migrations postgres "postgres://user:password@localhost:6433/checkout?sslmode=disable" up
