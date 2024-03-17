package store

import (
	"gorm.io/gorm"
)

type GlobalStore struct {
	TodoStore *ToDoStore
}

func NewGlobalStore(db *gorm.DB) *GlobalStore {
	todoStore := NewToDoStore(db)

	return &GlobalStore{TodoStore: todoStore}
}
