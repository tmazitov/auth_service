package staff

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

func (s *Staff) MakeTokenPair(ctx context.Context, claims jwt.MapClaims) (*TokenPair, error) {

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

	if err = s.Jwt.SaveToken(ctx, AccessPrefix, pair.Access, s.AccessDuration); err != nil {
		return nil, err
	}

	if err = s.Jwt.SaveToken(ctx, RefreshPrefix, pair.Refresh, s.RefreshDuration); err != nil {
		return nil, err
	}

	return pair, nil
}
