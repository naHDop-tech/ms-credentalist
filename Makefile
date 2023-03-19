DB_NAME=credentialist
DB_PASS=1qaz2wsx
DB_USER=credo
DB_HOST=0.0.0.0
BD_PORT=5432

create_db:
	docker run -d \
    --name credentialist \
    -e POSTGRES_PASSWORD=1qaz2wsx \
    -e POSTGRES_DB=credentialist \
    -e POSTGRES_USER=credo \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v db:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres
test:
	go test -v -cover ./...
sqlc_init:
	sqlc init
sqlc:
	sqlc generate
migrate_up:
	goose -dir=./db/migrations postgres "host=$(DB_HOST) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" up
migrate_down:
	goose -dir=./db/migrations postgres "host=$(DB_HOST) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" down
migrate_file:
	goose create $(file_name) sql
start:
	go run cmd/main.go