package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

type CodeCheckInput struct {
	Token string `json:"token" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

type CodeCheckOutput staff.TokenPair

type CodeCheckHandler struct {
	service.HandlerCoreBehavior[
		CodeCheckInput,
		CodeCheckOutput,
	]
	st *staff.Staff
}

func (h *CodeCheckHandler) Handle(ctx *gin.Context) {

	var (
		err    error
		email  string
		auth   *staff.UserAuth
		method *staff.UserAuthMethod
		pair   *staff.TokenPair
		userId int
	)

	if email, err = h.st.Conductor.VerifyCode(ctx, h.Input.Token, h.Input.Code); err != nil {
		staff.ResponseByCode(ctx, http.StatusBadRequest)
		return
	}

	auth = &staff.UserAuth{Email: email}
	method = &staff.UserAuthMethod{AuthMethodId: staff.EmailAuthMethod}
	if userId, err = h.st.Storage.UpdateUserAuthMethod(ctx, auth, method); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if err = h.st.Conductor.RemoveCode(ctx, h.Input.Token); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if pair, err = h.st.MakeTokenPair(ctx, staff.UserClaims(userId)); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	h.Output.Access = pair.Access
	h.Output.Refresh = pair.Refresh
}
