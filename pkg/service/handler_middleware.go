package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerCoreBehavior[I, O]) readInputMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h.ReadInput(c); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": http.StatusText(400)})
			return
		}
		c.Next()
	}
}

func (h *HandlerCoreBehavior[I, O]) writeOutputMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var status int = 200
		
		if ctx.Writer.Status() != 0 {
			status = ctx.Writer.Status()
		}

		ctx.JSON(status, h.Output)
		ctx.Next()
	}
}
