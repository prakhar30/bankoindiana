dockerpostgres:
	docker pull postgres

postgres:
	docker run --name postgresLatest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=something_secret -d postgres

createdb:
	docker exec -it postgresLatest createdb --username=root --owner=root banko_indiana

dropdb:
	docker exec -it postgresLatest dropdb banko_indiana

migrateup:
	migrate -path db/migration -database "postgresql://root:something_secret@localhost:5432/banko_indiana?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:something_secret@localhost:5432/banko_indiana?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen --package mockdb --destination db/mock/store.go github.com/prakhar30/bankoindiana/db/sqlc Store

.PHONY: dockerpostgres postgres createdb dropdb migrateup migratedown sqlc test server mock