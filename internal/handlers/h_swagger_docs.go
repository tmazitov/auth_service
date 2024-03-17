package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerDocsHandler struct {
}

func (h *SwaggerDocsHandler) BeforeMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

func (h *SwaggerDocsHandler) CoreBeforeMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

func (h *SwaggerDocsHandler) CoreAfterMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

func (h *SwaggerDocsHandler) AfterMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		ginSwagger.WrapHandler(swaggerfiles.Handler),
	}
}

func (h *SwaggerDocsHandler) Handle(c *gin.Context) {

}
