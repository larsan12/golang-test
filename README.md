# Домашнее задание

- Создать скелеты трёх сервисов по описанию АПИ из файла contracts.md
- Структуру проекта сделать с учетом разбиения на слои, бизнес-логику писать отвязанной от реализаций клиентов и хендлеров
- Все хендлеры отвечают просто заглушками
- Сделать удобный враппер для сервера по тому принципу, по которому делали на воркшопе
- Придумать самостоятельно удобный враппер для клиента
- Все межсервисные вызовы выполняются. Если хендлер по описанию из contracts.md должен ходить в другой сервис, он должен у вас это успешно делать в коде.
- Общение сервисов по http-json-rpc
- должны успешно проходить make precommit и make run-all в корневой папке
- Наладить общение с product-service (в хендлере Checkout.listCart). Токен для общения с product-service получить, написав в личку @pav5000

## urls

grafana http://localhost:3000/
jaeger http://localhost:16686/search

## how to run:

```
make up-db
make run-services
cd loms && make migration-run
cd checkout && make migration-run
```

## checkout config

```
token: QudSiFFeXkqUFEs7fDTxgLAn
services:
  loms: "loms:50052"
  product: "route256.pavl.uk:8082"
port: "50051"
db: "postgres://user:password@checkout-pgbouncer:6433/checkout"
productServiceRateLiming: 10
getProductPoolAmount: 5
kafkaBrokers:
  - "kafka:29092"
kafkaTopic: "test"
tracesUrl: "http://localhost:14268/api/traces"

```

## loms config

```
port: "50052"
db: "postgres://user:password@loms-pgbouncer:6432/loms"
OrderExpirationTime: 10
kafkaBrokers:
  - "kafka:29092"
kafkaTopic: "test"
tracesUrl: "http://localhost:14268/api/traces"
```
