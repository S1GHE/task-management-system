package users

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	id       uuid.UUID
	email    string
	password string
	created  time.Time
}

func NewUser(id uuid.UUID, email string, password string, created time.Time) (*User, error) {
	return &User{
		id:       id,
		email:    email,
		password: password,
		created:  created,
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

func (u *User) Created() time.Time {
	return u.created
}
