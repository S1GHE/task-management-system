package projects

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	id           uuid.UUID
	name         string
	description  string
	status       ProjectStatus
	dateCreate   time.Time
	dateComplete time.Time
	tasks        []Task
}

func NewProject(id uuid.UUID,
	status ProjectStatus,
	name, description string,
	dateCreate, dateComplete time.Time,
	tasks []Task) (*Project, error) {
	return &Project{
		id:           id,
		name:         name,
		description:  description,
		status:       status,
		dateCreate:   dateCreate,
		dateComplete: dateComplete,
		tasks:        tasks,
	}, nil
}

func (p *Project) ID() uuid.UUID {
	return p.id
}

func (p *Project) Name() string {
	return p.name
}

func (p *Project) Description() string {
	return p.description
}

func (p *Project) Status() ProjectStatus {
	return p.status
}

func (p *Project) DateCreate() time.Time {
	return p.dateCreate
}

func (p *Project) DateComplete() time.Time {
	return p.dateComplete
}

func (p *Project) Tasks() []Task {
	return p.tasks
}
