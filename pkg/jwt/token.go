package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *JwtStorage) CreateToken(ctx context.Context, claims jwt.MapClaims, duration time.Duration) (string, error) {
	claims["exp"] = time.Now().Add(duration).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(s.config.private)
}

func (s *JwtStorage) VerifyToken(ctx context.Context, name string, token string) (jwt.MapClaims, error) {

	if err := s.isExists(ctx, name, token); err != nil {
		return nil, ErrInvalidToken
	}

	return s.verifyToken(ctx, token)
}

func (s *JwtStorage) RemoveToken(ctx context.Context, name string, token string) error {
	return s.redis.Del(ctx, name+token).Err()
}

func (s *JwtStorage) SaveToken(ctx context.Context, name string, token string, duration time.Duration) error {
	return s.redis.Set(ctx, name+token, "", duration).Err()
}
