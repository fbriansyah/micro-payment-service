DB_URL=postgresql://root:secret@localhost:5432/db_payment?sslmode=disable

rabbit-mq:
	docker run --name rabbit-mq -p 5672:5672 -d rabbitmq:3.9-alpine

postgres:
	docker run --name pg-local -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pg-local createdb --username=root --owner=root db_payment

dropdb:
	docker exec -it pg-local dropdb db_payment

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

re-db: dropdb createdb migrateup

sqlc-win:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

run:
	go run ./cmd/main.go

build-image:
	docker build -t efner/payment-microservice:1.0 .

deploy: build-image
	docker push efner/payment-microservice:1.0

.PHONY: postgres createdb migrateup migrateup1 migratedown migratedown1 new_migration re-db run build-image
