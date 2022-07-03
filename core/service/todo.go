package service

import (
	"github.com/panutat-p/hexagonal-todo-gin/core/domain"
	"github.com/panutat-p/hexagonal-todo-gin/core/port"
	"net/http"
)

type TodoHandler struct {
	store port.Storer
}

func NewTodoHandler(store port.Storer) *TodoHandler {
	return &TodoHandler{
		store: store,
	}
}

func (h *TodoHandler) CreateNewTask(c port.Context) {
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
