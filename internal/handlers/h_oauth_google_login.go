package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

type OauthGoogleLoginHandler struct {
	service.HandlerMonoWriteBehavior[struct {
		Url string `json:"url"`
	}]
	st *staff.Staff
}

func (h *OauthGoogleLoginHandler) Handle(ctx *gin.Context) {

	var (
		state string
		url string
	)

	state = generateOauthState()

	ctx.SetCookie("oauth-state", state, 3600, "", "", false, true)

	url = h.st.Config.Google.AuthCodeURL(state)
	ctx.Redirect(http.StatusSeeOther, url)
	h.Output.Url = url
}

func generateOauthState() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}
