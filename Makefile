
build-all:
	cd checkout && GOOS=linux make build
	cd loms && GOOS=linux make build
	cd notifications && GOOS=linux make build

run-services: build-all
	sudo docker compose up -d loms --force-recreate --build
	sudo docker compose up -d checkout --force-recreate --build

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit

up-loms-db:
	sudo docker compose up -d loms-db --build
	sudo docker compose up -d loms-pgbouncer --build

up-checkout-db:
	sudo docker compose up -d checkout-db --build
	sudo docker compose up -d checkout-pgbouncer --build

up-db: up-loms-db up-checkout-db

run-loms:
	sudo docker compose up -d loms --force-recreate --build

run-checkout:
	sudo docker compose up -d checkout --force-recreate --build

down: 
	sudo docker compose down

clean-volumes:
	sudo docker compose down --volumes