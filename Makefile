
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
	sudo docker compose up -d loms-db --force-recreate --build
	sudo docker compose up -d loms-pgbouncer --force-recreate --build

up-checkout-db:
	sudo docker compose up -d checkout-db --force-recreate --build
	sudo docker compose up -d checkout-pgbouncer --force-recreate --build

up-db: up-loms-db up-checkout-db

down: 
	sudo docker compose down

clean-volumes:
	sudo docker compose down --volumes