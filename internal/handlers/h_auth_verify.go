package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tmazitov/auth_service.git/internal/staff"
)

type SendCodeHandler struct {
	st    *staff.Staff
	input struct {
		Email string `json:"email" binding:"required,email"`
	}
	output struct {
		Token string `json:"token"`
	}
}

func (h SendCodeHandler) Handle(ctx *gin.Context) {

	var err error

	if err = ctx.ShouldBindJSON(&h.input); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	if h.output.Token, err = h.st.Conductor.SendCode(h.input.Email); err != nil {
		ctx.JSON(500, gin.H{
			"error": "internal error",
		})
		return
	}

	ctx.JSON(200, h.output)
}

func (h SendCodeHandler) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
