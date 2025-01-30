package projects

type ProjectStatus string

const (
	ProjectStatusActive    ProjectStatus = "active"
	ProjectStatusCompleted ProjectStatus = "completed"
)

type TaskStatus string

const (
	TasksStatusLow     TaskStatus = "low"
	TasksStatusAverage TaskStatus = "average"
	TasksStatusHigh    TaskStatus = "high"
)
