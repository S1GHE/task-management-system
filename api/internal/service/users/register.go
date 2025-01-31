package users

import (
	"net/http"
	"task-management-system-api/pkg/generators"
	"task-management-system-api/pkg/request"
)

func (u *UserService) Register(req request.CreateUser) (error, int) {
	req.Password = generators.GenerateHashPassword(req.Password)

	if err := u.repo.CreateUser(req); err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
