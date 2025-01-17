package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	log *logrus.Logger
}

func New(log *logrus.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) InitHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /api/tasks", NewGetTasks(h.log))

	h.log.Info("Handler init")
	return mux
}
