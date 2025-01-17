package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	log    *logrus.Logger
}

func New(port string, handler http.Handler, log *logrus.Logger) *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		log: log,
	}
}

func (s *Server) Run() error {
	s.log.Infof("Start server on port %s", s.server.Addr)
	return s.server.ListenAndServe()
}
