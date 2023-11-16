package service

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

type Claims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("ErrNoMetadataInContext")
	}

	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", fmt.Errorf("ErrNoAuthorizationInMetadata")
	}

	return token[0], nil
}
