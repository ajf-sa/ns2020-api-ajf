createdb:
	docker exec -it db01 createdb --username=admin --owner=admin simple_api

dropdb:
	docker exec -it db01 dropdb simple_api -U admin

psql:
	docker exec -it db01 psql simple_api -U admin

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir db/migration -seq $$name

migratever:
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5432/simple_api?sslmode=disable" -verbose version

migrate:
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5432/simple_api?sslmode=disable" -verbose up


rollback:
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5432/simple_api?sslmode=disable" -verbose down

migrateto:
	@read -p "Enter migration version: " version; \
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5432/simple_api?sslmode=disable" -verbose up $$version

rollbackto:
	@read -p "Enter migration version: " version; \
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5432/simple_api?sslmode=disable" -verbose goto $$version

migrateforce:
	@read -p "Enter migration version: " version; \
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5432/simple_api?sslmode=disable" -verbose force $$version

sqlc:
	sqlc generate

run:
	go run main.go

build :
	go build -o .

echo :
	echo $<
	
.PHONY: build run createdb dropdb psql migrate rollback sqlc migration echo migratever migrateto rollbackto migrateforce