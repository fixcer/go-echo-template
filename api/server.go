package api

import (
	"github.com/labstack/echo/v4"
)

// Server serves HTTP requests for our banking usecase.
type Server struct {
	router *echo.Echo
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(router *echo.Echo) *Server {
	return &Server{
		router: router,
	}
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Start(address)
}
