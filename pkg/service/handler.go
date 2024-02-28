package service

import "github.com/gin-gonic/gin"

type Handler interface {
	Handle(ctx *gin.Context)
	Middleware() []gin.HandlerFunc
}

type Endpoint struct {
	Method  string
	Path    string
	Handler Handler
}
