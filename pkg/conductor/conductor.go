package cond

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tmazitov/auth_service.git/pkg/jwt"
)

type ConductorConfig struct {
	SenderPort       int    `json:"senderPort"  binding:"required"`
	SenderPass       string `json:"senderPass"  binding:"required"`
	SenderEmail      string `json:"senderEmail" binding:"required"`
	MailTitle        string `json:"mailTitle"`
	MailCodeLength   int    `json:"mailCodeLength"`
	MailCodeDuration int    `json:"mailCodeDuration"`
	MailTemplatePath string `json:"mailTemplatePath"  binding:"required"`
	TokenSecret      string `json:"tokenSecret"  binding:"required"`
	CodeRefreshDelay int    `json:"codeRefreshDelay:"`
}

type Conductor struct {
	emailChan       chan messageInfo
	mailTemplate    string
	redis           *redis.Client
	jwt             *jwt.JwtStorage
	config          *ConductorConfig
	mailDuration    time.Duration
	refreshDuration time.Duration
}

func validateConfig(config *ConductorConfig) error {
	if config == nil || config.MailTemplatePath == "" {
		return ErrInvalidParams
	}

	if config.TokenSecret == "" {
		return ErrInvalidParams
	}

	if config.MailCodeLength == 0 {
		config.MailCodeLength = DefaultCodeLength
	}

	if config.MailCodeDuration == 0 {
		config.MailCodeDuration = DefaultCodeDuration
	}

	if config.MailTitle == "" {
		config.MailTitle = "Auth Code"
	}

	if config.CodeRefreshDelay == 0 {
		config.CodeRefreshDelay = DefaultCodeRefreshDelay
	}

	return nil
}

func NewConductor(redis *redis.Client, config *ConductorConfig) (*Conductor, error) {

	var (
		conductor *Conductor
		storage   *jwt.JwtStorage
		template  string
		err       error
	)

	if err = validateConfig(config); err != nil {
		return nil, err
	}

	if template, err = openHTMLTemplate(config.MailTemplatePath); err != nil {
		return nil, err
	}

	if storage, err = jwt.NewJwtStorage([]byte(config.TokenSecret), redis); err != nil {
		return nil, err
	}

	conductor = &Conductor{
		emailChan:       make(chan messageInfo),
		jwt:             storage,
		mailTemplate:    template,
		redis:           redis,
		config:          config,
		mailDuration:    time.Duration(config.MailCodeDuration) * time.Minute,
		refreshDuration: time.Duration(config.CodeRefreshDelay) * time.Minute,
	}

	conductor.start()

	return conductor, nil
}

func (c *Conductor) start() {

	var emailChan chan messageInfo = c.emailChan

	go func() {
		c.worker(emailChan)
	}()
}
