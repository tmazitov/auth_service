package staff

import "github.com/golang-jwt/jwt/v5"

const (
	AccessPrefix  string = "access:"
	RefreshPrefix string = "refresh:"
)

func UserClaims(id int) jwt.MapClaims {

	var claims jwt.MapClaims = jwt.MapClaims{}

	claims["id"] = id

	return claims
}
