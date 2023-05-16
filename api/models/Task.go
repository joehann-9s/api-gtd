package models

import "gorm.io/gorm"

type TaskType string
const (
	SimpleTask TaskType = "simple task"
	ProjectTask TaskType = "project"
)

type TaskState int
const (
	Inbox TaskState = iota
	Timer
	Hotbed
	Main
	Done
	Archived
)


type Task struct {
	gorm.Model
	Title       string
	Description string
	Type        TaskType
	State       TaskState
	Category    []string
	UserID      uint
	SubTasks    []Task `gorm:"foreignKey:ParentTaskID"`
	ParentTask  *Task  `gorm:"foreignKey:ParentTaskID"`
}
