package todo

import "github.com/gin-gonic/gin"

// NewGinHandler convert HandlerAdapter to GinHandler
// because gin router accept GinHandler only
func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}

// MyContext composition
// wrapper of GinContext
type MyContext struct {
	*gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{
		Context: c,
	}
}

func (c *MyContext) Bind(v interface{}) error {
	//return c.ShouldBindJSON(v)
	return c.Context.ShouldBindJSON(v)
}

func (c *MyContext) Json(statusCode int, v interface{}) {
	c.Context.JSON(statusCode, v)
}

func (c *MyContext) TransactionId() string {
	return c.Request.Header.Get("Transaction-ID")
}

func (c *MyContext) Audience() string {
	if aud, ok := c.Get("aud"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}
