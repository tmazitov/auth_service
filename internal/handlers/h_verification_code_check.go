package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	)

	if email, err = h.st.Conductor.VerifyCode(ctx, h.Input.Token, h.Input.Code); err != nil {
		staff.ResponseByCode(ctx, http.StatusBadRequest)
		return
	}

	auth = &staff.UserAuth{Email: email}
	method = &staff.UserAuthMethod{AuthMethodId: staff.EmailAuthMethod}
	if err = h.st.Storage.AddUserAuthMethod(ctx, auth, method); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if err = h.st.Conductor.RemoveCode(ctx, h.Input.Token); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if err = h.makeTokenPair(ctx); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}
}

func (h *CodeCheckHandler) makeTokenPair(ctx *gin.Context) error {

	var (
		err    error
		claims jwt.MapClaims
	)
	claims = staff.UserClaims(1)

	if h.Output.Access, err = h.st.Jwt.CreateToken(ctx, claims, h.st.AccessDuration); err != nil {
		return err
	}

	if h.Output.Refresh, err = h.st.Jwt.CreateToken(ctx, claims, h.st.AccessDuration); err != nil {
		return err
	}

	if err = h.st.Jwt.SaveToken(ctx, staff.AccessPrefix, h.Output.Access, h.st.AccessDuration); err != nil {
		return err
	}

	if err = h.st.Jwt.SaveToken(ctx, staff.RefreshPrefix, h.Output.Refresh, h.st.RefreshDuration); err != nil {
		return err
	}

	return nil
}
