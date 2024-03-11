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
	return c.redis.Set(ctx, codePrefix+token, code, c.mailDuration).Err()
}

func (c *Conductor) SendCode(ctx context.Context, email string, ip string) (string, error) {

	var (
		err        error
		code       string
		token      string
		blocker    string
		claims     jwt.MapClaims = newClaims(email)
		clientInfo string        = fmt.Sprintf("%s:%s", email, ip)
	)

	if blocker, err = c.getRefreshBlocker(ctx, clientInfo); err != nil {
		return "", err
	}

	if blocker != "" {
		return "", ErrCodeRefreshBlock
	}

	code = generateCode(c.config.MailCodeLength)

	c.emailChan <- messageInfo{email: email, code: code}

	if token, err = c.jwt.CreateToken(ctx, claims, c.mailDuration); err != nil {
		return "", err
	}

	if err = c.saveCode(ctx, token, code); err != nil {
		return "", err
	}

	if err = c.setRefreshBlocker(ctx, clientInfo, token); err != nil {
		return "", err
	}

	return token, nil
}

func (c *Conductor) VerifyCode(ctx context.Context, token string, code string) (string, error) {

	var (
		err       error
		cmd       *redis.StringCmd
		email     string
		claims    jwt.MapClaims
		savedCode string
	)

	if claims, err = c.jwt.VerifyToken(ctx, codePrefix, token); err != nil {
		return "", err
	}

	cmd = c.redis.Get(ctx, codePrefix+token)
	if savedCode, err = cmd.Result(); err != nil {
		return "", err
	}

	if savedCode != code {
		return "", ErrInvalidCode
	}

	if claims["email"] == nil {
		return "", ErrInvalidCode
	}

	email = claims["email"].(string)
	if email == "" {
		return "", ErrInvalidCode
	}

	return email, nil
}

func (c *Conductor) RemoveCode(ctx context.Context, token string) error {
	return c.jwt.RemoveToken(ctx, codePrefix, token)
}

func (c *Conductor) setRefreshBlocker(ctx context.Context, clientInfo string, token string) error {
	return c.redis.Set(ctx, codeRefresh+clientInfo, token, c.refreshDuration).Err()
}

func (c *Conductor) getRefreshBlocker(ctx context.Context, clientInfo string) (string, error) {
	var (
		cmd    *redis.StringCmd = c.redis.Get(ctx, codeRefresh+clientInfo)
		err    error
		result string
	)

	if result, err = cmd.Result(); err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return result, nil
}
