package cond

import "github.com/golang-jwt/jwt/v5"

const (
	prefix string = "cond-verification:"
)

func newClaims(email string) *jwt.MapClaims {

	var claims jwt.MapClaims = jwt.MapClaims{}

	claims["email"] = email

	return &claims
}
