postgres:
	docker run --name postgres-new -p 5433:5432   -e POSTGRES_PASSWORD=secret -d postgres:15.1-alpine

createdb:
	docker exec -it postgres-new createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres-new  dropdb simple_bank -U postgres

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up 

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

mockstore:
	mockgen --build_flags=--mod=mod  -package mockdb -destination db/mock/store.go github.com/maan19/bank-app-go/db/sqlc  Store

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc mockstore test server
