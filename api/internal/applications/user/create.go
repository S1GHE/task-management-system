package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-management-system-api/pkg/request"
)

func (h *UsersHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req request.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := h.validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	response, err := json.Marshal(map[string]string{
		"msg": fmt.Sprintf("Пользователь %s успшно создан", req.Login),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
