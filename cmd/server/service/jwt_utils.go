package service

import (
	"context"
	"proteitestcase/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	tokenDutation = 30 * time.Minute
)

type AuthToken struct {
	Token string
}

// RequireTransportSecurity implements credentials.PerRPCCredentials.
func (*AuthToken) RequireTransportSecurity() bool {
	return true
}

func IsCorrectPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(tokenDutation).Unix(),
		"login": login,
	})

	secretKey, err := config.GetSecretKey()
	if err != nil {
		return "", err
	}
	return token.SignedString([]byte(secretKey))
}

func (t AuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": t.Token,
	}, nil
}
