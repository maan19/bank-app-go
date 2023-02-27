package gapi

import (
	"fmt"

	db "github.com/maan19/bank-app-go/db/sqlc"
	"github.com/maan19/bank-app-go/pb"
	"github.com/maan19/bank-app-go/token"
	"github.com/maan19/bank-app-go/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

// Creates a new HTTP server and creates routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create paseto maker: %w", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
