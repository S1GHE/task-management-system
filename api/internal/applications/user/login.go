package user

import (
	"encoding/json"
	"net/http"
)

func (h *UsersHandler) LoginUser(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(map[string]string{"msg": "Авторизация"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
