package services

import (
	"github.com/panutat-p/hexagonal-todo-gin/todo/domain"
	"github.com/panutat-p/hexagonal-todo-gin/todo/ports"
	"net/http"
)

type Handler struct {
	store ports.Storer
}

func NewHandler(store ports.Storer) *Handler {
	return &Handler{
		store: store,
	}
}

type Context interface {
	Bind(interface{}) error
	Json(int, interface{})
	TransactionId() string
	Audience() string
}

func (h *Handler) CreateNewTask(c Context) {
	var todo domain.Todo
	if err := c.Bind(&todo); err != nil {
		c.Json(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err := h.store.Create(&todo)
	if err != nil {
		c.Json(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.Json(http.StatusCreated, map[string]interface{}{
		"id": todo.ID,
	})
}
