package store

import (
	"github.com/cuongnd9/go-grpc/model"
	"gorm.io/gorm"
)

type ToDoStore struct {
	db *gorm.DB
}

func NewToDoStore(db *gorm.DB) *ToDoStore {
	return &ToDoStore{db: db}
}

func (s *ToDoStore) Create(input *model.Todo) *uint32 {
	err := s.db.Create(&input).Error
	if err != nil {
		return nil
	}
	return &input.ID
}

func (s *ToDoStore) Update(id uint32, title string, description string) uint32 {
	err := s.db.Model(&model.Todo{}).Where("id = ?", id).Updates(map[string]string{"title": title, "description": description})
	if err != nil {
		return 0
	}
	return 1
}

func (s *ToDoStore) Delete(id uint32) uint32 {
	err := s.db.Model(&model.Todo{}).Where("id = ?", id).Delete(&model.Todo{}).Error
	if err != nil {
		return 0
	}
	return 1
}

func (s *ToDoStore) Read(id uint32) *model.Todo {
	var result = &model.Todo{}
	err := s.db.Model(&model.Todo{}).First(result, id).Error
	if err != nil {
		return nil
	}
	return result
}

func (s *ToDoStore) ReadAll() []model.Todo {
	var result = []model.Todo{}
	s.db.Model(&model.Todo{}).Find(result)
	return result
}
