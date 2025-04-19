package cond

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tmazitov/auth_service.git/pkg/conductor/messages"
	"github.com/tmazitov/auth_service.git/pkg/jwt"
)

type Conductor struct {
	emailChan       chan *messages.MessageInfo
	redis           *redis.Client
	jwt             *jwt.JwtStorage
	config          *ConductorConfig
	mailDuration    time.Duration
	refreshDuration time.Duration
}

func NewConductor(redis *redis.Client, config *ConductorConfig) (*Conductor, error) {

	var (
		conductor *Conductor
		storage   *jwt.JwtStorage
		err       error
	)

	if err = config.Validate(); err != nil {
		return nil, err
	}
	if storage, err = jwt.NewJwtStorage([]byte(config.TokenSecret), redis); err != nil {
		return nil, err
	}

	conductor = &Conductor{
		emailChan:       make(chan *messages.MessageInfo),
		jwt:             storage,
		redis:           redis,
		config:          config,
		mailDuration:    time.Duration(config.MailCodeDuration) * time.Minute,
		refreshDuration: time.Duration(config.CodeRefreshDelay) * time.Minute,
	}

	conductor.start()

	return conductor, nil
}

func (c *Conductor) start() {

	var emailChan chan *messages.MessageInfo = c.emailChan

	go func() {
		c.worker(emailChan)
	}()
}
