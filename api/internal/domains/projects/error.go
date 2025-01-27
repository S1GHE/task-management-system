package projects

import "errors"

var (
	ErrTaskNotFound = errors.New("tasks not found")
)

var (
	ErrProjectNotFound = errors.New("project not found")
)
