package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskType string

const (
	SimpleTask  TaskType = "simple task"
	ProjectTask TaskType = "project"
)

type TaskState int

const (
	Inbox TaskState = iota
	Reminder
	Hotbed
	Main
	Done
	Archived
)

type Task struct {
	gorm.Model
	Title        string `gorm:"not null"`
	Description  string
	UserID       uint
	ParentID     uint
	SubTasks     []Task `gorm:"foreignKey:ParentID"`
	Type         TaskType
	State        TaskState
	Categories   []Category `gorm:"many2many:task_categories;"`
	ReminderDate *time.Time
}
