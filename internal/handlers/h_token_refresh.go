package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

type TokenRefreshHandler struct {
	service.HandlerCoreBehavior[
		struct {
			Refresh string `json:"refresh_token" binding:"required"`
		},
		staff.TokenPair,
	]
	st *staff.Staff
}

func (h *TokenRefreshHandler) Handle(ctx *gin.Context) {

}
