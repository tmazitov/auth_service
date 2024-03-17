package staff

import "github.com/golang-jwt/jwt/v5"

const (
	AccessPrefix  string = "access:"
	RefreshPrefix string = "refresh:"
)

type OauthUserInfo struct {
	Email      string `json:"email"`
	FirstName  string `json:"given_name"`
	LastName   string `json:"family_name"`
	IsVerified bool   `json:"verified_email"`
	Picture    string `json:"picture"`
}

func UserClaims(id int) jwt.MapClaims {

	var claims jwt.MapClaims = jwt.MapClaims{}

	claims["id"] = id

	return claims
}
