package users

import "errors"

var (
	ErrUserNotFound        = errors.New("users not found")
	ErrUserInvalidPassword = errors.New("invalid password")
)
