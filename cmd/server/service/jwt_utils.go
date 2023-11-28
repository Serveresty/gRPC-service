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
	debugLg.Debug().Str("action", "Checking is correct password")
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	debugLg.Debug().Str("action", "Generating hash from password")
	return string(bytes), err
}

func CreateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(tokenDutation).Unix(),
		"login": login,
	})

	secretKey, err := config.GetSecretKey()
	if err != nil {
		lg.Err(err).Msg("Error while getting secret key")
		return "", err
	}
	debugLg.Debug().Str("action", "Created token")
	return token.SignedString([]byte(secretKey))
}

func (t AuthToken) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	debugLg.Debug().Str("action", "Get request metadata")
	return map[string]string{
		"authorization": t.Token,
	}, nil
}
