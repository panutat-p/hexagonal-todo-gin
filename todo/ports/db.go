package ports

import (
	"github.com/panutat-p/hexagonal-todo-gin/todo/domain"
)

type Storer interface {
	Create(todo *domain.Todo) error
	Query(title string) error
}
