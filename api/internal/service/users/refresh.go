package users

import (
	"net/http"
	"task-management-system-api/pkg/generators"
)

func (u *UserService) Refresh(token string) (string, error, int) {
	userID, err := u.repo.FindByToken(token)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	accessToken, err := generators.GenerateAccessToken(userID)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	return accessToken, nil, http.StatusOK
}
