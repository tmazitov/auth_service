package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_service "github.com/tmazitov/auth_service.git/internal/proto/user_service"
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

// @Summary Check verification code
// @Description Check the verification code for a given token
// @Tags Standard Auth
// @Accept json
// @Produce json
// @Param input body CodeCheckInput true "Input parameters"
// @Success 200 {object} CodeCheckOutput "Verification code check successful"
// @Failure 400 {object} staff.ErrorResponse "Bad request"
// @Failure 500 {object} staff.ErrorResponse "Internal server error"
// @Router /verification/code/check [post]
func (h *CodeCheckHandler) Handle(ctx *gin.Context) {

	var (
		err        error
		email      string
		auth       *staff.UserAuth
		method     *staff.UserAuthMethod
		pair       *staff.TokenPair
		respUserID *user_service.GetUserIDResponse
		respCreate *user_service.CreateUserResponse
		userId     int
	)

	if email, err = h.st.Conductor.VerifyCode(ctx, h.Input.Token, h.Input.Code); err != nil {
		staff.ResponseByCode(ctx, http.StatusBadRequest)
		return
	}

	auth = &staff.UserAuth{Email: email}
	method = &staff.UserAuthMethod{AuthMethodId: staff.EmailAuthMethod}
	if _, err = h.st.Storage.UpdateUserAuthMethod(ctx, auth, method); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if err = h.st.Conductor.RemoveCode(ctx, h.Input.Token); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	if respUserID, err = h.st.UserService.GetUserID(ctx, &user_service.GetUserIDRequest{Email: email}); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}
	userId = int(respUserID.UserId)
	if userId == 0 {
		if respCreate, err = h.st.UserService.CreateUser(ctx, &user_service.CreateUserRequest{Email: email}); err != nil {
			staff.ResponseByError(ctx, err)
			return
		}
		userId = int(respCreate.UserId)
	}

	if pair, err = h.st.MakeTokenPair(ctx, staff.UserClaims(userId)); err != nil {
		staff.ResponseByError(ctx, err)
		return
	}

	h.Output.Access = pair.Access
	h.Output.Refresh = pair.Refresh
}
