package model

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title       string     `gorm:"column:title"`
	Description string     `gorm:"column:description"`
	Reminder    *time.Time `gorm:"column:reminder"`
}

// TableName get table name
func (todo *Todo) TableName() string {
	return "todo"
}
