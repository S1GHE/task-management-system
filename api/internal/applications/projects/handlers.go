package projects

import "github.com/go-playground/validator/v10"

type ProjectHandler struct {
	validator *validator.Validate
}

func SetupHandler(validator *validator.Validate) *ProjectHandler {
	return &ProjectHandler{
		validator: validator,
	}
}
