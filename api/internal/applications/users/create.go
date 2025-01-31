package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-management-system-api/pkg/request"
)

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req request.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err := h.validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err, statusCode := h.service.Register(req)
	if err != nil {
		http.Error(w, err.Error(), statusCode)

		return
	}

	response, err := json.Marshal(map[string]string{
		"msg": fmt.Sprintf("Пользователь %s успшно создан", req.Email),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
