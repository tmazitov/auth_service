package cond

import (
	"strings"

	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
)

type ConductorConfig struct {
	SenderPort       int    `json:"senderPort"  binding:"required"`
	SenderPass       string `json:"senderPass"  binding:"required"`
	SenderEmail      string `json:"senderEmail" binding:"required"`
	MailTitle        string `json:"mailTitle"`
	MailCodeLength   int    `json:"mailCodeLength"`
	MailCodeDuration int    `json:"mailCodeDuration"`
	MailTemplatePath string `json:"mailTemplatePath"  binding:"required"`
}

type Conductor struct {
	emailChan    chan gomail.Message
	errChan      chan error
	mailTemplate string
	redis        *redis.Client
	config       *ConductorConfig
}

func validateConfig(config *ConductorConfig) error {
	if config == nil || config.MailTemplatePath == "" {
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

	return nil
}

func NewConductor(redis *redis.Client, config *ConductorConfig) (*Conductor, error) {

	var (
		conductor *Conductor
		template  string
		err       error
	)

	if err = validateConfig(config); err != nil {
		return nil, err
	}

	if template, err = openHTMLTemplate(config.MailTemplatePath); err != nil {
		return nil, err
	}

	conductor = &Conductor{
		emailChan:    make(chan gomail.Message),
		errChan:      make(chan error),
		mailTemplate: template,
		redis:        redis,
		config:       config,
	}

	conductor.start()

	return conductor, nil
}

func (c *Conductor) SendCode(email string) (string, error) {
	var (
		code    string
		text    string
		message *gomail.Message
	)

	code = generateCode(c.config.MailCodeLength)
	text = strings.Replace(c.mailTemplate, "{{.VerificationCode}}", code, 1)
	message = gomail.NewMessage()
	message.SetHeader("From", c.config.SenderEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", c.config.MailTitle)
	message.SetBody("text/html", text)

	c.emailChan <- *message

	return code, nil
}

func (c *Conductor) start() {

	var emailChan chan gomail.Message = c.emailChan

	go func() {
		c.worker(emailChan)
	}()
}

func (c *Conductor) VerifyCode(token string, code string) bool {
	return false
}
