package main

import (
	"database/sql"
	"log"

	"github.com/maan19/bank-app-go/api"
	db "github.com/maan19/bank-app-go/db/sqlc"
	"github.com/maan19/bank-app-go/util"
)

func main() {
	config, err := util.Loadconfig(".")
	if err != nil {
		log.Fatal("ERROR loading configs", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error creating db:", err)
	}

	store := db.NewSQLStore(conn)
	server := api.NewServer(store)

	server.Start(config.ServerAddress)

}
