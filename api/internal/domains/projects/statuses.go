package projects

type ProjectStatus string

const (
	ProjectStatusActive    ProjectStatus = "active"
	ProjectStatusCompleted ProjectStatus = "completed"
)

type TaskStatus string

const (
	TasksStausLow     TaskStatus = "low"
	TasksStausAverage TaskStatus = "average"
	TasksStausHigh    TaskStatus = "high"
)
