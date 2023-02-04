postgres:
	docker run --name postgres21 -p 5433:5432   -e POSTGRES_PASSWORD=secret -d postgres:15.1-alpine

createdb:
	docker exec -it postgres21 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres21  dropdb simple_bank -U postgres

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown
