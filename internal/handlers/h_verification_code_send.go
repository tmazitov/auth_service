package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/auth_service.git/internal/staff"
	cond "github.com/tmazitov/auth_service.git/pkg/conductor"
)

type CodeSendHandler struct {
	st    *staff.Staff
	input struct {
		Email string `json:"email" binding:"required,email"`
	}
	output struct {
		Token string `json:"token"`
	}
}

func (h *CodeSendHandler) Handle(ctx *gin.Context) {

	var err error

	if err = ctx.ShouldBindJSON(&h.input); err != nil {
		staff.ResponseByCode(ctx, http.StatusBadRequest)
		return
	}

	h.output.Token, err = h.st.Conductor.SendCode(ctx, h.input.Email, ctx.ClientIP())
	if err == cond.ErrCodeRefreshBlock {
		staff.ResponseByCode(ctx, http.StatusForbidden)
		return
	}

	if err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	ctx.JSON(200, h.output)
}

func (h *CodeSendHandler) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}