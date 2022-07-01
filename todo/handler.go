package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	store Storer
}

func NewHandler(store Storer) *Handler {
	return &Handler{
		store: store,
	}
}

// Context adapter spec for API
type Context interface {
	Bind(interface{}) error
	Json(int, interface{})
	TransactionId() string
	Audience() string
}

// Storer adapter spec for SPI
type Storer interface {
	Create(*Todo) error
	Query(*Todo) error
}

func (h *Handler) CreateNewTask(c Context) {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.Json(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.store.Create(&todo)
	if err != nil {
		c.Json(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Json(http.StatusCreated, gin.H{
		"id": todo.ID,
	})
}
