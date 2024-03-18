package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

type ServiceDocsHandler struct {
	service.HandlerClearBehavior
}

func (h *ServiceDocsHandler) Handle(c *gin.Context) {
}

func (h *ServiceDocsHandler) AfterMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		ginSwagger.WrapHandler(swaggerfiles.Handler),
	}
}
