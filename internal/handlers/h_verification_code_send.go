package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/auth_service.git/internal/staff"
	cond "github.com/tmazitov/auth_service.git/pkg/conductor"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

type CodeSendInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CodeSendOutput struct {
	Token string `json:"token"`
}

type CodeSendHandler struct {
	service.HandlerCoreBehavior[
		CodeSendInput,
		CodeSendOutput,
	]
	st *staff.Staff
}

func (h *CodeSendHandler) Handle(ctx *gin.Context) {

	var err error

	h.Output.Token, err = h.st.Conductor.SendCode(ctx, h.Input.Email, ctx.ClientIP())
	if err == cond.ErrCodeRefreshBlock {
		staff.ResponseByCode(ctx, http.StatusForbidden)
		return
	}

	if err != nil {
		staff.ResponseByError(ctx, err)
		return
	}
}
