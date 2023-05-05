package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	"github.com/Mikatech/sso-auth-server/go/pkg/config"
	"github.com/Mikatech/sso-auth-server/go/pkg/database"
	"github.com/golang-jwt/jwt"
)

func Generate(user database.User) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	addClaims(token, user)
	return "", nil
}

// TODO: Change Id per ID
func addClaims(token *jwt.Token, user database.User) {
	token.Claims = jwt.MapClaims{
		"preferred_username": user.Username,
		"email":              user.Email,
		"email_verified":     true,
		"role":               user.Role,
		"sub":                user.Id,
		"iss":                config.Config.Issuer,
		"exp":                time.Now().Add(time.Duration(config.Config.ExpireIn) * time.Minute),
		"iat":                time.Now(),
		"alg":                "RS256",
	}
}

func loadPrivateKey() (*rsa.PrivateKey, error) {
	keyString, err := os.ReadFile(config.Config.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyString)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
