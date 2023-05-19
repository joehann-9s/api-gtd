package models

import "gorm.io/gorm"

type SubTask struct {
	gorm.Model
	Title       string
	Description string
	ParentID    uint
}
