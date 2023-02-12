DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

run_container:
	docker run --name postgresBank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest postgres -N 1000

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

migrate_up1:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up 1

migrate_down:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down

migrate_down1:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down 1

sqlc_cmd:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover -short -count=1 ./...

test_special:
	go test .\internal\db\sqlc\ -timeout 30s -run ^TestTransferTx -v -count=1

cmd_write_raw_queries:
	docker exec -it postgresBank psql -U root simple_bank

server:
	go run main.go

.PHONY: run_container start_container create_db drop_db migrate_up migrate_down migrate_up1 migrate_down1 sqlc_cmd test test_special cmd_write_raw_queries server


# Notes:
# for disable cache in test using flag: -count=1
# postgres -N 1000 means max_connections for concurrency