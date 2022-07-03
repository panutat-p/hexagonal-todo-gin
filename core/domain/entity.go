package domain

import (
	"time"
)

type Todo struct {
	Title     string `json:"title"`
	ID        uint   `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName
// Gorm, provide custom table name
func (Todo) TableName() string {
	return "todolist"
}
