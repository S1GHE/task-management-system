package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GetTasks struct {
	log *logrus.Logger
}

func NewGetTasks(log *logrus.Logger) *GetTasks {
	return &GetTasks{
		log: log,
	}
}

func (h *GetTasks) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(map[string]string{"msg": "hello"})

	if err != nil {
		h.log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(data); err != nil {
		h.log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
