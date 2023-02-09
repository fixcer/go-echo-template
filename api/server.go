package api

import (
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking usecase.
type Server struct {
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
