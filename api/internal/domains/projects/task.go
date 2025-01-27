package projects

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	id           uuid.UUID
	name         string
	description  string
	status       TaskStatus
	dateCreate   time.Time
	dateComplete time.Time
}

func NewTask(id uuid.UUID, status TaskStatus, name, description string, dateCreate, dateComplete time.Time) (*Task, error) {
	return &Task{
		id:           id,
		name:         name,
		description:  description,
		status:       status,
		dateCreate:   dateCreate,
		dateComplete: dateComplete,
	}, nil
}

func (t *Task) ID() uuid.UUID {
	return t.id
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() TaskStatus {
	return t.status
}

func (t *Task) DateCreate() time.Time {
	return t.dateCreate
}

func (t *Task) DateComplete() time.Time {
	return t.dateComplete
}
