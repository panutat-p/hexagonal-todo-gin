package repository

import (
	"github.com/panutat-p/hexagonal-todo-gin/core/domain"
	"gorm.io/gorm"
)

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db}
}

type GormStore struct {
	*gorm.DB
}

func (gs GormStore) CreateItem(todo *domain.Todo) error {
	return gs.Create(todo).Error
}

func (gs GormStore) QueryItems(title string) error {
	//TODO implement me
	panic("🟥 Not Implemented Yet")
}
