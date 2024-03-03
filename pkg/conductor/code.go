package cond

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func generateCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < length; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}

func (c *Conductor) saveCode(ctx context.Context, token string, code string) error {
	return c.redis.Set(ctx, prefix+token, code, c.mailDuration).Err()
}

func (c *Conductor) SendCode(ctx context.Context, email string) (string, error) {

	var (
		err    error
		code   string
		token  string
		claims jwt.MapClaims = *newClaims(email)
	)

	code = generateCode(c.config.MailCodeLength)

	c.emailChan <- messageInfo{email: email, code: code}

	claims["email"] = email

	if token, err = c.jwt.CreateToken(ctx, claims, c.mailDuration); err != nil {
		return "", err
	}

	if err = c.saveCode(ctx, token, code); err != nil {
		return "", err
	}

	return token, nil
}

func (c *Conductor) VerifyCode(ctx context.Context, token string, code string) error {

	var (
		err       error
		cmd       *redis.StringCmd
		savedCode string
	)

	if _, err = c.jwt.VerifyToken(ctx, prefix, token); err != nil {
		return err
	}

	cmd = c.redis.Get(ctx, prefix+token)
	if savedCode, err = cmd.Result(); err != nil {
		return err
	}

	if savedCode != code {
		return ErrInvalidCode
	}

	return nil
}

func (c *Conductor) RemoveCode(ctx context.Context, token string) error {
	return c.jwt.RemoveToken(ctx, prefix, token)
}
