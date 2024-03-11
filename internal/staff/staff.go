package staff

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tmazitov/auth_service.git/internal/config"
	cond "github.com/tmazitov/auth_service.git/pkg/conductor"
	"github.com/tmazitov/auth_service.git/pkg/jwt"
)

type Staff struct {
	Conductor       *cond.Conductor
	Jwt             *jwt.JwtStorage
	Config          *config.Config
	AccessDuration  time.Duration
	RefreshDuration time.Duration
	Storage         IStorage
}

func NewStaff(config *config.Config) *Staff {

	var (
		accessDuration  time.Duration = time.Duration(config.JwtAccess) * time.Minute
		refreshDuration time.Duration = time.Duration(config.JwtRefresh) * 24 * time.Hour
	)

	return &Staff{
		Config:          config,
		Conductor:       nil,
		Jwt:             nil,
		AccessDuration:  accessDuration,
		RefreshDuration: refreshDuration,
	}
}

func (s *Staff) SetJwt(redis *redis.Client, secret string) error {
	var err error

	if s.Jwt, err = jwt.NewJwtStorage([]byte(secret), redis); err != nil {
		return err
	}
	return nil
}

func (s *Staff) SetConductor(redis *redis.Client, conf *cond.ConductorConfig) error {
	var err error
	s.Conductor, err = cond.NewConductor(redis, conf)
	if err != nil {
		return err
	}
	return nil
}

func (s *Staff) SetStorage(storage IStorage) {
	s.Storage = storage
}
