package users

import (
	"encoding/json"
	"net/http"
	"task-management-system-api/pkg/request"
)

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req request.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokens, err, statusCode := h.service.Login(req)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	response, err := json.Marshal(map[string]string{
		"access_Token":  tokens.AccessToken,
		"refresh_Token": tokens.RefreshToken,
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
