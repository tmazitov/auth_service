package staff

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseByData(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, data)
}

func ResponseByCode(ctx *gin.Context, code int) {
	ctx.JSON(code, gin.H{
		"message": http.StatusText(code),
	})
}

func ResponseByError(ctx *gin.Context, err error) {
	ctx.Errors = append(ctx.Errors, &gin.Error{Err: err, Type: gin.ErrorTypePublic})
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server error",
	})
	panic(err)
}
