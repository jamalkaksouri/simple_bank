package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jamalkaskouri/simple_bank/internal/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.PUT("/accounts", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}