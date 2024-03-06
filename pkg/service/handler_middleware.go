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
	return func(c *gin.Context) {
		c.JSON(200, h.Output)
		c.Next()
	}
}
