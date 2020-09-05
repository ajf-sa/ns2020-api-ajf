create first migrate:
    migrate create -ext sql -dir db/migration -seq init_schema


init sqlc:
    sqlc init
