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
	Reminder
	Hotbed
	Main
	Done
	Archived
)

type Task struct {
    gorm.Model
    Title         string
    Description   string
    Type          TaskType
    State         TaskState
    UserID        uint
    SubTasks      []Task `gorm:"foreignKey:ParentTaskID"`
    ParentTask    *Task  `gorm:"foreignKey:ParentTaskID"`
    ReminderDate  *time.Time
    Category	  []string
}
