package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/panutat-p/hexagonal-todo-gin/core/port"
)

// NewGinHandler
// convert WrappedHandler to GinHandler
func NewGinHandler(h port.WrappedHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(newMyContext(c))
	}
}

type MyContext struct {
	context *gin.Context
}

func newMyContext(c *gin.Context) *MyContext {
	return &MyContext{
		context: c,
	}
}

func (c *MyContext) Bind(v interface{}) error {
	//return c.ShouldBindJSON(v)
	return c.context.ShouldBindJSON(v)
}

func (c *MyContext) Json(statusCode int, v interface{}) {
	c.context.JSON(statusCode, v)
}

func (c *MyContext) TransactionId() string {
	return c.context.Request.Header.Get("Transaction-ID")
}

func (c *MyContext) Audience() string {
	if aud, ok := c.context.Get("aud"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}
