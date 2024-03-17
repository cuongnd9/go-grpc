package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          uint32 `gorm:"column:id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
}

// TableName get table name
func (todo *Todo) TableName() string {
	return "todo"
}
