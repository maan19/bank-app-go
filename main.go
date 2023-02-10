package main

import (
	"database/sql"
	"log"

	"github.com/maan19/bank-app-go/api"
	db "github.com/maan19/bank-app-go/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:secret@localhost:5433/simple_bank?sslmode=disable"
	address  = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error creating db:", err)
	}

	store := db.NewSQLStore(conn)
	server := api.NewServer(store)

	server.Start(address)

}
