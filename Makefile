DB_SOURCE= "postgresql://postgres:secret@localhost:5433/simple_bank?sslmode=disable"
postgres:
	docker run --name postgres-new --network bankapp-network -p 5433:5432   -e POSTGRES_PASSWORD=secret -d postgres:15.1-alpine

createdb:
	docker exec -it postgres-new createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres-new  dropdb simple_bank -U postgres

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 

migrateup1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

dbdocs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

mock:
	mockgen --build_flags=--mod=mod  -package mockdb -destination db/mock/store.go github.com/maan19/bank-app-go/db/sqlc  Store

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=bank_app \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc 

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7.0.0-alpine


.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc mock test server dbdocs db_schema proto evans redis
