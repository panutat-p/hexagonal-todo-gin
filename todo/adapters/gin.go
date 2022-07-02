package adapters

import (
	"github.com/gin-gonic/gin"
	"github.com/panutat-p/hexagonal-todo-gin/todo/services"
)

func NewGinHandler(handler func(services.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}

type MyContext struct {
	context *gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
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
