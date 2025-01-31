package users

import (
	userInfra "task-management-system-api/internal/infrastructure/users"
	"task-management-system-api/internal/service"
)

type UserService struct {
	service.UserService
	repo *userInfra.PostgresRepo
}

func NewUserService(repo *userInfra.PostgresRepo) *UserService {
	return &UserService{repo: repo}
}
