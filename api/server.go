package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vitaLemoTea/myBank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

//NewServer create a new Http server and  set up routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccount)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
