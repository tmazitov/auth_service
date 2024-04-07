package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

type PingInput struct {
	Text string `json:"text" binding:"required"`
}

type PingOutput struct {
	Message string `json:"message"`
}

type PingHandler struct {
	service.HandlerMonoWriteBehavior[PingOutput]
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func (h *PingHandler) Handle(ctx *gin.Context) {
	h.Output.Message = "pong"
}
