package router

import (
	"github.com/gin-gonic/gin"
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
// func(ctx port.Context) is Handler
func (r *GinRouter) POST(path string, h port.Handler) {
	r.Engine.POST(path, NewGinHandler(h))
}
