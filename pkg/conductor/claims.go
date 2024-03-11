package cond

import "github.com/golang-jwt/jwt/v5"

const (
	codePrefix  string = "cond-verification:"
	codeRefresh string = "cond-verification-refresh:"
)

func newClaims(email string) jwt.MapClaims {

	var claims jwt.MapClaims = jwt.MapClaims{}

	claims["email"] = email

	return claims
}
