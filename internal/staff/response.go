package staff

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func ResponseByData(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, data)
}

func ResponseByCode(ctx *gin.Context, code int) {

	ctx.Status(code)
	if code >= 400 {
		ctx.AbortWithStatus(code)
	}
	ctx.JSON(code, ErrorResponse{Message: http.StatusText(code)})
}

func ResponseByError(ctx *gin.Context, err error) {
	var code int = http.StatusInternalServerError
	ctx.Errors = append(ctx.Errors, &gin.Error{Err: err, Type: gin.ErrorTypePublic})
	ctx.JSON(code, ErrorResponse{Message: http.StatusText(code)})
	panic(err)
}
