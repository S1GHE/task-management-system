package projects

import (
	"encoding/json"
	"net/http"
)

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(map[string]string{"msg": "Создание проекта"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) CreateProjectTask(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(map[string]string{"msg": "Создание задачи связанной с проектом"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
