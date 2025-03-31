package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	addr   string
	router *gin.Engine
}

func NewServer(addr string) *Server {
	return &Server{
		addr:   addr,
		router: gin.Default(),
	}
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) Start() error {
	fmt.Printf("Starting server on %s\n", s.addr)
	return http.ListenAndServe(s.addr, s.router)
}
