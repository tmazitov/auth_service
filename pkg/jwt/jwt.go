package jwt

import (
	"github.com/redis/go-redis/v9"
)

type JwtConfig struct {
	private []byte
}

type JwtStorage struct {
	redis  *redis.Client
	config JwtConfig
}

func NewJwtStorage(private []byte, redis *redis.Client) (*JwtStorage, error) {
	return &JwtStorage{
		redis:  redis,
		config: JwtConfig{private: private},
	}, nil
}
