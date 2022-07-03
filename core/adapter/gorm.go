package adapter

import (
	"github.com/panutat-p/hexagonal-todo-gin/core/domain"
	"gorm.io/gorm"
)

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{
		db: db,
	}
}

type GormStore struct {
	db *gorm.DB
}

func (gs GormStore) Create(todo *domain.Todo) error {
	return gs.db.Create(todo).Error
}

func (gs GormStore) Query(title string) error {
	//TODO implement me
	panic("🟥 Not Implemented Yet")
}
