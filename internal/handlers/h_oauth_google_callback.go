package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/pkg/service"
	"golang.org/x/oauth2"
)

type OauthGoogleCallbackInput struct {
	Code  string `form:"code" binding:"required"`
	State string `form:"state" binding:"required"`
}

type OauthGoogleCallbackHandler struct {
	service.HandlerCoreBehavior[
		OauthGoogleCallbackInput,
		staff.TokenPair,
	]
	st *staff.Staff
}

// @Summary Google OAuth Callback
// @Description Handles the callback from Google OAuth authentication
// @Tags Google OAuth
// @Accept json
// @Produce json
// @Param code query string true "Authorization code received from Google"
// @Param state query string true "State parameter received from Google"
// @Success 200 {object} staff.TokenPair "Token pair containing access and refresh tokens"
// @Failure 400 {object} staff.ErrorResponse "Bad request"
// @Failure 403 {object} staff.ErrorResponse "Forbidden"
// @Failure 500 {object} staff.ErrorResponse "Internal server error"
// @Router /oauth/google/callback [get]
func (h *OauthGoogleCallbackHandler) Handle(ctx *gin.Context) {

	var (
		userId      int
		user        *staff.OauthUserInfo
		userData    []byte
		userClaims  jwt.MapClaims
		token       *oauth2.Token
		pair        *staff.TokenPair
		err         error
		cookieState string
	)

	if cookieState, err = ctx.Cookie("oauth-state"); err != nil {
		staff.ResponseByCode(ctx, 403)
		return
	}

	if h.Input.State != cookieState {
		staff.ResponseByCode(ctx, 400)
		return
	}

	ctx.SetCookie("oauth-state", "", -1, "", "", false, true)

	if token, err = h.st.Config.Google.Exchange(ctx, h.Input.Code); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if userData, err = h.st.OauthUserData(token); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if err = json.Unmarshal(userData, &user); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	auth := &staff.UserAuth{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	method := &staff.UserAuthMethod{
		AuthMethodId: staff.GoogleAuthMethod,
	}

	if userId, err = h.st.Storage.UpdateUserAuthMethod(ctx, auth, method); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	userClaims = staff.UserClaims(userId)

	if pair, err = h.st.MakeTokenPair(ctx, userClaims); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	h.Output = *pair
}
