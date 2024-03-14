package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

type TokenRefreshInput struct {
	Refresh string `json:"refresh_token" binding:"required"`
}

type TokenRefreshHandler struct {
	service.HandlerCoreBehavior[
		TokenRefreshInput,
		staff.TokenPair,
	]
	st *staff.Staff
}

func (h *TokenRefreshHandler) Handle(ctx *gin.Context) {

	var (
		err     error
		oldPair *staff.TokenPair
		newPair *staff.TokenPair
		claims  jwt.MapClaims
		access  string = h.st.GetAccessToken(ctx)
	)

	if err = h.st.Jwt.IsExists(ctx, staff.AccessPrefix, access); err != nil {
		staff.ResponseByCode(ctx, 401)
		return
	}

	if claims, err = h.st.GetRefreshClaims(ctx, h.Input.Refresh); err != nil {
		staff.ResponseByCode(ctx, 401)
		return
	}

	oldPair = &staff.TokenPair{Access: access, Refresh: h.Input.Refresh}

	if err = h.st.RemoveTokenPair(ctx, oldPair); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if newPair, err = h.st.MakeTokenPair(ctx, claims); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	h.Output.Access = newPair.Access
	h.Output.Refresh = newPair.Refresh
}
