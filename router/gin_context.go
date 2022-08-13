package router

import (
	"github.com/gin-gonic/gin"
)

// GinContext is a wrapper of real gin context
type GinContext struct {
	*gin.Context
}

func (c *GinContext) Bind(v interface{}) error {
	//return c.ShouldBindJSON(v)
	return c.ShouldBindJSON(v)
}

func (c *GinContext) Json(statusCode int, v interface{}) {
	c.JSON(statusCode, v)
}

func (c *GinContext) TransactionId() string {
	return c.Request.Header.Get("Transaction-ID")
}

func (c *GinContext) Audience() string {
	if aud, ok := c.Get("aud"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}
