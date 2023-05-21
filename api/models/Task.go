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
	Inbox    TaskType = "inbox"
	Reminder TaskType = "reminder"
	Hotbed   TaskType = "hotbed"
	Main     TaskType = "main"
	Done     TaskType = "done"
	Archived TaskType = "archived"
)

type Task struct {
	gorm.Model
	Title        string `gorm:"not null"`
	Description  string
	UserID       uint
	Type         TaskType
	State        TaskState
	Categories   []*Category `gorm:"many2many:task_categories;"`
	ReminderDate *time.Time
	SubTasks     []*SubTask `gorm:"foreignkey:ParentID"`
}
