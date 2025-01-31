package users

import (
	"github.com/go-playground/validator/v10"
	usersInfra "task-management-system-api/internal/infrastructure/users"
	"task-management-system-api/internal/service/users"
)

type UserHandler struct {
	validator *validator.Validate
	service   *users.UserService
}

func SetupHandlers(validator *validator.Validate, repo *usersInfra.PostgresRepo) *UserHandler {
	service := users.NewUserService(repo)

	return &UserHandler{
		validator: validator,
		service:   service,
	}
}
