LOMS_MIGRATION_DSN = "postgres://user:password@localhost:5433/loms?sslmode=disable"
CHECKOUT_MIGRATION_DSN = "postgres://user:password@localhost:5432/checkout?sslmode=disable"

debug: up-db migrate

build-all:
	cd checkout && GOOS=linux make build
	cd loms && GOOS=linux make build
	cd notifications && GOOS=linux make build

run-services: build-all migrate
	sudo docker compose up -d loms --force-recreate --build
	sudo docker compose up -d checkout --force-recreate --build

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit

up-loms-db:
	sudo docker compose up -d loms-db --build

up-checkout-db:
	sudo docker compose up -d checkout-db --build

up-db: up-loms-db up-checkout-db

run-loms:
	sudo docker compose up -d loms --force-recreate --build

run-checkout:
	sudo docker compose up -d checkout --force-recreate --build

down: 
	sudo docker compose down

clean-volumes:
	sudo docker compose down --volumes

loms-migrate:
	goose -dir ./loms/migrations postgres ${LOMS_MIGRATION_DSN} up

checkout-migrate:
	goose -dir ./checkout/migrations postgres ${CHECKOUT_MIGRATION_DSN} up

migrate: loms-migrate checkout-migrate

up-kafka:
	sudo docker compose up -d zookeeper
	sudo docker compose up -d kafka

up-metrics:
	sudo docker compose up -d prometheus --build
	sudo docker compose up -d grafana --build
	sudo docker compose up -d jaeger --build

run-all: build-all
	sudo docker compose up -d --force-recreate --build
	make migrate