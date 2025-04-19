package cond

import "fmt"

type AMQPConfig struct {
	Host string `json:"host"  binding:"required"`
	Port int    `json:"port"  binding:"required"`
	User string `json:"user"  binding:"required"`
	Pass string `json:"pass"  binding:"required"`
}

func (c *AMQPConfig) Validate() error {

	if c == nil || c.Port == 0 || c.Host == "" {
		return ErrInvalidAMQPConfig
	}

	if c.User == "" || c.Pass == "" {
		return ErrInvalidAMQPConfig
	}

	return nil
}

func (c *AMQPConfig) GetDSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d", c.User, c.Pass, c.Host, c.Port)
}

type ConductorConfig struct {
	AMQPConfig       AMQPConfig `json:"amqp"  binding:"required"`
	MailCodeLength   int        `json:"mailCodeLength"`
	MailCodeDuration int        `json:"mailCodeDuration"`
	TokenSecret      string     `json:"tokenSecret"  binding:"required"`
	CodeRefreshDelay int        `json:"codeRefreshDelay:"`
}

func (c *ConductorConfig) Validate() error {

	if c == nil {
		return ErrInvalidParams
	}

	if err := c.AMQPConfig.Validate(); err != nil {
		return err
	}

	if c.TokenSecret == "" {
		return ErrInvalidParams
	}

	if c.MailCodeLength == 0 {
		c.MailCodeLength = DefaultCodeLength
	}

	if c.MailCodeDuration == 0 {
		c.MailCodeDuration = DefaultCodeDuration
	}

	if c.CodeRefreshDelay == 0 {
		c.CodeRefreshDelay = DefaultCodeRefreshDelay
	}

	return nil
}
