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

migrateup1:
	migrate -path db/migration -database "postgresql://root:something_secret@localhost:5432/banko_indiana?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:something_secret@localhost:5432/banko_indiana?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:something_secret@localhost:5432/banko_indiana?sslmode=disable" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen --package mockdb --destination db/mock/store.go github.com/prakhar30/bankoindiana/db/sqlc Store

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans: 
	evans --host localhost --port 9091 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7.4-alpine

redis-ping:
	docker exec -it redis redis-cli ping

.PHONY: dockerpostgres postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock proto evans redis redis-ping new_migration