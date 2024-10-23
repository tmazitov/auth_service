package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/service/handler"
)

type TokenRefreshInput struct {
	Refresh string `json:"refreshToken" binding:"required"`
}

type TokenRefreshHandler struct {
	handler.CoreBehavior[
		TokenRefreshInput,
		staff.TokenPair,
	]
	st *staff.Staff
}

// @Summary Refreshes the token pair
// @Description Refreshes the token pair using the provided refresh token
// @Tags Standard Auth
// @Accept json
// @Produce json
// @Param refresh_token body TokenRefreshInput true "Refresh Token"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} staff.TokenPair "Token Pair"
// @Failure 401 {object} staff.ErrorResponse "Unauthorized"
// @Failure 500 {object} staff.ErrorResponse "Internal Server Error"
// @Router /token/refresh [post]
// @Security ApiKeyAuth
func (h *TokenRefreshHandler) Handle(ctx *gin.Context) {
	var (
		err     error
		oldPair *staff.TokenPair
		newPair *staff.TokenPair
		claims  jwt.MapClaims
		access  string = h.st.GetAccessToken(ctx)
	)

	if err = h.st.Jwt.IsExists(ctx, staff.AccessPrefix, access); err != nil {
		staff.ResponseByCode(ctx, http.StatusUnauthorized)
		return
	}

	if claims, err = h.st.GetRefreshClaims(ctx, h.Input.Refresh); err != nil {
		staff.ResponseByCode(ctx, http.StatusUnauthorized)
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
