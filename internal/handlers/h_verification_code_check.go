package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tmazitov/auth_service.git/internal/staff"
)

type CodeCheckHandler struct {
	st    *staff.Staff
	input struct {
		Token string `json:"token" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}
	output struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}
}

func (h *CodeCheckHandler) Handle(ctx *gin.Context) {

	var err error

	if err = ctx.ShouldBindJSON(&h.input); err != nil {
		staff.ResponseByCode(ctx, http.StatusBadRequest)
		return
	}

	if err = h.st.Conductor.VerifyCode(ctx, h.input.Token, h.input.Code); err != nil {
		staff.ResponseByCode(ctx, http.StatusBadRequest)
		return
	}

	if err = h.st.Conductor.RemoveCode(ctx, h.input.Token); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if err = h.makeTokenPair(ctx); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	staff.ResponseByData(ctx, 200, h.output)
}

func (h *CodeCheckHandler) makeTokenPair(ctx *gin.Context) error {

	var (
		err    error
		claims jwt.MapClaims
	)
	claims = staff.UserClaims(1)

	if h.output.Access, err = h.st.Jwt.CreateToken(ctx, claims, h.st.AccessDuration); err != nil {
		return err
	}

	if h.output.Refresh, err = h.st.Jwt.CreateToken(ctx, claims, h.st.AccessDuration); err != nil {
		return err
	}

	if err = h.st.Jwt.SaveToken(ctx, staff.AccessPrefix, h.output.Access, h.st.AccessDuration); err != nil {
		return err
	}

	if err = h.st.Jwt.SaveToken(ctx, staff.RefreshPrefix, h.output.Refresh, h.st.RefreshDuration); err != nil {
		return err
	}

	return nil
}

func (h *CodeCheckHandler) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
