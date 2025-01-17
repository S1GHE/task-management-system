package handler

import "net/http"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) InitHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /api/tasks", NewGetTasks())

	return mux
}
