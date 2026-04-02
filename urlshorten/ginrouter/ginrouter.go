package ginrouter

import (
	"github.com/gin-gonic/gin"
	"github.com/pallat/urlshorten/httprouter"
)

func NewHandler(handler func(httprouter.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := NewContext(ctx)
		handler(c)
	}
}

type context struct {
	ctx *gin.Context
}

func NewContext(c *gin.Context) *context {
	return &context{ctx: c}
}

func (c *context) Param(k string) string {
	return c.ctx.Param(k)
}

func (c *context) Bind(v any) error {
	return c.ctx.Bind(v)
}

func (c *context) Status(code int) {
	c.ctx.Status(code)
}

func (c *context) JSON(code int, v any) {
	c.ctx.JSON(code, v)
}

func (c *context) Redirect(code int, url string) {
	c.ctx.Redirect(code, url)
}
