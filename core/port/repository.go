package port

import (
	"github.com/panutat-p/hexagonal-todo-gin/core/domain"
)

type Storer interface {
	Create(todo *domain.Todo) error
	Query(title string) error
}
