package users

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	email    string
	password string
}

func NewUser(id uuid.UUID, email string, password string) (*User, error) {
	return &User{
		id:       id,
		email:    email,
		password: password,
	}, nil
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}
