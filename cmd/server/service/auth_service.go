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
		lg.Err(err).Msg("Error while getting auth data")
		return &api.LoginResponce{Token: ""}, err
	}
	debugLg.Debug().Str("action", "Get auth data success")

	hashPass, err := HashPassword(password)
	if err != nil {
		lg.Err(err).Msg("Error while getting hash password")
		return &api.LoginResponce{Token: ""}, err
	}
	debugLg.Debug().Str("action", "Hash pass success")

	if login == in.Login && IsCorrectPassword(hashPass, in.Password) {
		token, err := CreateToken(in.Login)
		if err != nil {
			lg.Err(err).Msg("Error while creating token")
			return nil, fmt.Errorf("Error while creating token")
		}
		debugLg.Debug().Str("action", "Login success")
		return &api.LoginResponce{Token: token}, nil
	}
	lg.Err(err).Msg("Bad credentials")
	return &api.LoginResponce{Token: ""}, fmt.Errorf("Bad credentials")
}

func CheckAuth(ctx context.Context) (string, error) {
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		lg.Err(err).Msg("Error while getting token from context")
		return "", fmt.Errorf("get token from context error: %v", err)
	}
	debugLg.Debug().Str("action", "Get token from context success")

	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			lg.Error().Msg("Invalid algorithm")
			return "", fmt.Errorf("Error Invalid Algorithm: %v", err)
		}
		secretKey, err1 := config.GetSecretKey()
		if err1 != nil {
			lg.Err(err1).Msg("Error while getting secret key")
			return nil, err1
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		lg.Err(err).Msg("JWT parse error")
		return "", fmt.Errorf("jwt parse error: %v", err)
	}

	if !token.Valid {
		lg.Err(err).Msg("Error invalid token")
		return "", fmt.Errorf("Err Invalid Token: %v", err)
	}

	debugLg.Debug().Str("action", "Check auth success")
	return clientClaims.Login, nil
}
