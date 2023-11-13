package service

import (
	"context"
	clientservice "proteitestcase/cmd/client/service"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	userStore  clientservice.UserStore
	jwtManager *clientservice.JWTManager
	api.UnimplementedAuthServiceServer
}

func NewAuthServer(userStore clientservice.UserStore, jwtManager *clientservice.JWTManager) *AuthServer {
	return &AuthServer{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
}

func (server *AuthServer) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponce, error) {
	user, err := server.userStore.Find(req.GetLogin())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect user data")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while generate access token")
	}

	res := &api.LoginResponce{Token: token}

	return res, nil
}
