package store

import "database/sql"

type GlobalStore struct {
	TodoStore *ToDoStore
}

func NewGlobalStore(db *sql.DB) *GlobalStore {
	todoStore := NewToDoStore(db)

	return &GlobalStore{TodoStore: todoStore}
}
