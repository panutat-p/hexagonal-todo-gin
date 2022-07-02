package adapters

import (
	"github.com/panutat-p/hexagonal-todo-gin/todo/domain"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{
		db: db,
	}
}

func (gs *GormStore) Create(todo *domain.Todo) error {
	return gs.db.Create(todo).Error
}

func (gs *GormStore) Query(title string) error {
	//TODO implement me
	panic("🟥 Not Implemented Yet")
}
