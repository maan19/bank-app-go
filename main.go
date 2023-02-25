package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/maan19/bank-app-go/api"
	db "github.com/maan19/bank-app-go/db/sqlc"
	"github.com/maan19/bank-app-go/gapi"
	"github.com/maan19/bank-app-go/pb"
	"github.com/maan19/bank-app-go/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Error creating server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Error creating listener:", err)
	}

	log.Printf("starting grpc server on %s", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Error creating server:", err)
	}

	server.Start(config.HTTPServerAddress)
}
