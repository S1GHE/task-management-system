package service

import (
	"task-management-system-api/pkg/request"
	"task-management-system-api/pkg/response"
)

type UserService interface {
	Register(user request.CreateUser) (error, int)
	Login(user request.CreateUser) (*response.Login, error, int)
	Refresh(token string) (string, error, int)
}
