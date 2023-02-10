package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/maan19/bank-app-go/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

// Creates a new HTTP server and creates routing.
func NewServer(store *db.SQLStore) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	server.router = router
	return server
}

// Start runs the HTTP server on given address
func (server *Server) Start(address string) {
	server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
