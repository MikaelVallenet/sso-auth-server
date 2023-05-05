package token

import (
	"github.com/Mikatech/sso-auth-server/go/pkg/database"
	"github.com/golang-jwt/jwt"
)

func Generate(user database.User) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	return "", nil
}

func AddClaims(token *jwt.Token, user database.User) {
	token.Claims = jwt.MapClaims{
		"preferred_username": user.Username,
		"email":              user.Email,
		"email_verified":     true,
		"role":               user.Role,
	}
}
