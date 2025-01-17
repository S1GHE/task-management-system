package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	log.Printf("Сервер запущен на %s", s.server.Addr)
	return s.server.ListenAndServe()
}
