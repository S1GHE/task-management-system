package handler

import (
	"encoding/json"
	"net/http"
)

type GetTasks struct{}

func NewGetTasks() *GetTasks {
	return &GetTasks{}
}

func (h *GetTasks) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(map[string]string{"msg": "hello"})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
