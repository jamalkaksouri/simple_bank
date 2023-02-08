DB_URL=postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable

run_container:
	docker run --name postgresBank -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest postgres -N 1000
	#postgres -N 1000 means max_connections for concurrency

start_container:
	docker start postgresBank

create_db:
	docker exec -it postgresBank createdb --username=root --owner=root simple_bank

drop_db:
	docker exec -it postgresBank dropdb simple_bank

migrate_init:
	migrate create -ext sql -dir internal/db/migration -seq init_schema

migrate_up:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down

sqlc_cmd:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover -short ./...
	#go test -v -cover -count=1 ./...

test_special:
	go test .\internal\db\sqlc\ -timeout 30s -run ^TestTransferTx -v -count=1

cmd_write_raw_queries:
	docker exec -it postgresBank psql -U root simple_bank

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination internal/db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

.PHONY: run_container start_container create_db drop_db migrate_up migrate_down sqlc_cmd test test_special cmd_write_raw_queries server