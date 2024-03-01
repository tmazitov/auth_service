package staff

import (
	"github.com/redis/go-redis/v9"
	cond "github.com/tmazitov/auth_service.git/pkg/conductor"
)

type Staff struct {
	Conductor *cond.Conductor
}

func NewStaff() *Staff {
	return &Staff{
		Conductor: nil,
	}
}

func (s *Staff) SetConductor(redis *redis.Client, conf *cond.ConductorConfig) error {
	var err error
	s.Conductor, err = cond.NewConductor(redis, conf)
	if err != nil {
		return err
	}
	return nil
}
