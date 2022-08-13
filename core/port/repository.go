package port

import (
	"github.com/panutat-p/hexagonal-todo-gin/core/domain"
)

type Storer interface {
	CreateItem(todo *domain.Todo) error
	QueryItems(title string) error
}
