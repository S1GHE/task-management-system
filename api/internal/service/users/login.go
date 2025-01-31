package users

import (
	"net/http"
	userDomain "task-management-system-api/internal/domains/users"
	"task-management-system-api/pkg/generators"
	"task-management-system-api/pkg/request"
	"task-management-system-api/pkg/response"
)

func (u *UserService) Login(req request.CreateUser) (*response.Login, error, int) {
	user, err := u.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	if user.Password() != generators.GenerateHashPassword(req.Password) {
		return nil, userDomain.ErrUserInvalidPassword, http.StatusUnauthorized
	}

	accessToken, err := generators.GenerateAccessToken(user.ID())
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	refreshToken, err := generators.GenerateRefreshToken()
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	if err := u.repo.CreateToken(user.ID(), refreshToken); err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return &response.Login{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil, http.StatusOK
}
