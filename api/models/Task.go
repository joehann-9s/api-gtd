package models

import "gorm.io/gorm"

type TaskType int
type TaskState int
type TaskCategory int

const (
	SimpleTask TaskType = iota
	ProjectTask
)

const (
	Inbox TaskState = iota
	Timer
	Hotbed
	Main
	Done
	Archived
)

const (
	Job TaskCategory = iota
	UBO
	Education
	Leisure
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	Type        TaskType
	State       TaskState
	Category    TaskCategory
	UserID      uint
	SubTasks    []Task `gorm:"foreignKey:ParentTaskID"`
	ParentTask  *Task  `gorm:"foreignKey:ParentTaskID"`
}
