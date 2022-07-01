package todo

import "gorm.io/gorm"

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{
		db: db,
	}
}

func (gs *GormStore) Create(todo *Todo) error {
	return gs.db.Create(todo).Error
}

func (gs *GormStore) Query(todo *Todo) error {
	//TODO implement me
	panic("🟥 Not Implemented Yet")
}
