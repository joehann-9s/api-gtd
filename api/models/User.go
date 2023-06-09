package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string     `gorm:"not null"`
	LastName   string     `gorm:"not null"`
	Email      string     `gorm:"not null;unique_index"`
	Username   string     `gorm:"not null;unique_index"`
	Password   string     `gorm:"not null"`
	Tasks      []Task     `gorm:"foreignKey:UserID"`
	Categories []Category `gorm:"many2many:user_categories;"`
}
