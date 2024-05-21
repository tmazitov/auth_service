package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tmazitov/service/handler"
)

type ServiceDocsHandler struct {
	handler.DefaultBehavior
}

func (h *ServiceDocsHandler) Handle(ctx *gin.Context) {
}

func (h *ServiceDocsHandler) AfterMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		ginSwagger.WrapHandler(swaggerfiles.Handler),
	}
}
