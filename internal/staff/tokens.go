package staff

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
