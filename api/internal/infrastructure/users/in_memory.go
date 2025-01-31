package users

import "github.com/google/uuid"

// TODO --- Спросить а зачем в целом in_Memory?
type user struct {
	email    string
	password string
}

type InMemoryUserRepository struct {
	users map[uuid.UUID]user
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[uuid.UUID]user),
	}
}
