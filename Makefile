createdb:
	docker exec -it db01 createdb --username=admin --owner=admin simple_api

dropdb:
	docker exec -it db01 dropdb simple_api -U admin

psql:
	docker exec -it db01 psql simple_api -U admin

migrateup:
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5432/simple_api?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5432/simple_api?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
.PHONY: createdb dropdb psql migrateup migratedown sqlc