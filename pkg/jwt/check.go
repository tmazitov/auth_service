package jwt

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func (s *JwtStorage) verifyToken(ctx context.Context, token string) (jwt.MapClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return s.config.private, nil
	}

	jwtToken, err := jwt.Parse(token, keyFunc)
	if err != nil || !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	var ( 
		claims jwt.MapClaims	
		ok bool
	)

	if claims, ok = jwtToken.Claims.(jwt.MapClaims); !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
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
