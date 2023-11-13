package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"proteitestcase/cmd/server/service"
	"proteitestcase/internal/config"
	"proteitestcase/pkg/api"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crtFile = "./internal/server_data/openssl/server.crt"
	keyFile = "./internal/server_data/openssl/server.key"
)

const (
	secretKey     = "ultra-very-strong-secret-key"
	tokenDuration = 15 * time.Minute
)

func unaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func seedUser(userStore service.UserStore) error {
	login, password, err := config.GetAuthData()
	if err != nil {
		return err
	}
	return createUser(userStore, login, password)
}

func createUser(userStore service.UserStore, login string, password string) error {
	user, err := service.NewUser(login, password)
	if err != nil {
		return err
	}

	return userStore.Save(user)
}

func main() {
	if err := runServer(); err != nil {
		log.Fatal(err)
	}
}

func runServer() error {
	userStore := service.NewInMemUserStore()
	err := seedUser(userStore)
	if err != nil {
		return err
	}
	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)

	address, err1 := config.GetServerConnectionData()
	if err1 != nil {
		return err1
	}

	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	serverRegistrar := grpc.NewServer(opts...)

	api.RegisterAuthServiceServer(serverRegistrar, authServer)
	//api.RegisterDEMServer(serverRegistrar, &service.MyDEMServer{})

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	if err = serverRegistrar.Serve(listener); err != nil {
		fmt.Println("failed to serve: %s" + err.Error())
	}

	return nil
}
