package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/maan19/bank-app-go/db/sqlc"
	"github.com/maan19/bank-app-go/token"
	"github.com/maan19/bank-app-go/util"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
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

	server.addRoutes()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	return server, nil
}

// Start runs the HTTP server on given address
func (server *Server) Start(address string) {
	server.router.Run(address)
}

func (server *Server) addRoutes() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/token/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
