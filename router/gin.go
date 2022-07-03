package router

import (
	"github.com/gin-gonic/gin"
	"github.com/panutat-p/hexagonal-todo-gin/core/adapter"
	"github.com/panutat-p/hexagonal-todo-gin/core/port"
)

type GinRouter struct {
	*gin.Engine
}

func NewGinRouter() *GinRouter {
	r := gin.Default()
	// add middlewares
	//r.Use()
	return &GinRouter{
		r,
	}
}

// POST
// func(ctx port.Context) is WrappedHandler
func (r *GinRouter) POST(path string, h port.WrappedHandler) {
	r.Engine.POST(path, adapter.NewGinHandler(h))
}
