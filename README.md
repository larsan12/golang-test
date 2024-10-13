# цель
- Создать скелеты трёх сервисов по описанию АПИ из файла contracts.md
- Структуру проекта сделать с учетом разбиения на слои, бизнес-логику писать отвязанной от реализаций клиентов и хендлеров

## urls

grafana http://localhost:3000/
jaeger http://localhost:16686/search

## how to run:

```
make run-all
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
