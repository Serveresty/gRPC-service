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
		lg.Error().Msg("Err no metadata in context")
		return "", fmt.Errorf("ErrNoMetadataInContext")
	}

	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		lg.Error().Msg("Err no authorization in metadata")
		return "", fmt.Errorf("ErrNoAuthorizationInMetadata")
	}
	debugLg.Debug().Str("action", "Get token from context")
	return token[0], nil
}
