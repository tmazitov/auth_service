package staff

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

func (s *Staff) MakeTokenPair(ctx *gin.Context, claims jwt.MapClaims) (*TokenPair, error) {

	var (
		pair *TokenPair = &TokenPair{}
		err  error
	)

	if pair.Access, err = s.Jwt.CreateToken(ctx, claims, s.AccessDuration); err != nil {
		return nil, err
	}

	if pair.Refresh, err = s.Jwt.CreateToken(ctx, claims, s.RefreshDuration); err != nil {
		return nil, err
	}

	if err = s.Jwt.SaveToken(ctx, AccessPrefix, pair.Access, s.RefreshDuration); err != nil {
		return nil, err
	}

	if err = s.Jwt.SaveToken(ctx, RefreshPrefix, pair.Refresh, s.RefreshDuration); err != nil {
		return nil, err
	}

	return pair, nil
}

func (s *Staff) GetAccessToken(ctx *gin.Context) string {
	var authorizationValue string = ctx.GetHeader("Authorization")

	if authorizationValue == "" {
		return ""
	}

	if strings.Contains(authorizationValue, "Bearer") {
		return strings.Split(authorizationValue, " ")[1]
	}

	return authorizationValue
}

func (s *Staff) GetAccessClaims(ctx *gin.Context, token string) (jwt.MapClaims, error) {
	return s.Jwt.VerifyToken(ctx, AccessPrefix, token)
}

func (s *Staff) GetRefreshClaims(ctx *gin.Context, token string) (jwt.MapClaims, error) {
	return s.Jwt.VerifyToken(ctx, RefreshPrefix, token)
}

func (s *Staff) RemoveTokenPair(ctx *gin.Context, pair *TokenPair) error {
	var err error

	if err = s.Jwt.RemoveToken(ctx, AccessPrefix, pair.Access); err != nil {
		return err
	}
	if err = s.Jwt.RemoveToken(ctx, RefreshPrefix, pair.Refresh); err != nil {
		return err
	}
	return nil
}

func (s *Staff) OauthUserData(token *oauth2.Token) ([]byte, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
