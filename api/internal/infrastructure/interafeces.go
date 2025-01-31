package infrastructure

import (
	"github.com/google/uuid"
	"task-management-system-api/internal/domains/users"
	"task-management-system-api/pkg/request"
)

type UsersRepo interface {
	CreateUser(user *request.CreateUser) error
	CreateToken(userID uuid.UUID, token string) error
	GetUserByEmail(email string) (*users.User, error)
}
