package jwt

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func (s *JwtStorage) verifyToken(ctx context.Context, token string) (jwt.Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return s.config.private, nil
	}

	var claims jwt.MapClaims

	jwtToken, err := jwt.ParseWithClaims(token, claims, keyFunc)
	if err != nil || !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	return jwtToken.Claims, nil
}

func (s *JwtStorage) isExists(ctx context.Context, prefix string, token string) error {

	err := s.redis.Get(ctx, prefix+token).Err()
	if err == redis.Nil {
		return ErrTokenIsNotExist
	}

	if err != nil {
		return err
	}

	return nil
}
