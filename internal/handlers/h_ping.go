package handlers

import "github.com/gin-gonic/gin"

type PingHandler struct {
}

func (h PingHandler) Handle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h PingHandler) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
