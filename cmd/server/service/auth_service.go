package service

import (
	"context"
	"fmt"
	"proteitestcase/internal/config"
	"proteitestcase/pkg/api"

	"github.com/dgrijalva/jwt-go"
)

type AuthServer struct {
	api.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(_ context.Context, in *api.LoginRequest) (*api.LoginResponce, error) {
	login, password, err := config.GetAuthData()
	if err != nil {
		return &api.LoginResponce{Token: ""}, err
	}

	hashPass, err := HashPassword(password)
	if err != nil {
		return &api.LoginResponce{Token: ""}, err
	}

	if login == in.Login && IsCorrectPassword(hashPass, in.Password) {
		token, err := CreateToken(in.Login)
		if err != nil {
			return nil, fmt.Errorf("Error while creating token")
		}
		return &api.LoginResponce{Token: token}, nil
	}

	return &api.LoginResponce{Token: ""}, fmt.Errorf("Bad credentials")
}

func CheckAuth(ctx context.Context) (string, error) {
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("get token from context error: %v", err)
	}
	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			return "", fmt.Errorf("Error Invalid Algorithm: %v", err)
		}
		secretKey, err1 := config.GetSecretKey()
		if err1 != nil {
			return nil, err1
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("jwt parse error: %v", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("Err Invalid Token: %v", err)
	}

	return clientClaims.Login, nil
}
