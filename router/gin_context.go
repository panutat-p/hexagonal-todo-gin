package router

import (
	"github.com/gin-gonic/gin"
	"github.com/panutat-p/hexagonal-todo-gin/core/port"
)

// NewGinHandler convert Handler to real gin handler
func NewGinHandler(h port.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(newGinContext(c))
	}
}

// GinContext is a wrapper of real gin context
type GinContext struct {
	context *gin.Context
}

func newGinContext(c *gin.Context) *GinContext {
	return &GinContext{
		context: c,
	}
}

func (c *GinContext) Bind(v interface{}) error {
	//return c.ShouldBindJSON(v)
	return c.context.ShouldBindJSON(v)
}

func (c *GinContext) Json(statusCode int, v interface{}) {
	c.context.JSON(statusCode, v)
}

func (c *GinContext) TransactionId() string {
	return c.context.Request.Header.Get("Transaction-ID")
}

func (c *GinContext) Audience() string {
	if aud, ok := c.context.Get("aud"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}