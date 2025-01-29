package user

import "github.com/go-playground/validator/v10"

type UsersHandler struct {
	validator *validator.Validate
}

func SetupHandlers(validator *validator.Validate) *UsersHandler {
	return &UsersHandler{
		validator: validator,
	}
}
