package main

import (
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
	tokenDuration = 15 * time.Minute
)

func main() {
	if err := runServer(); err != nil {
		log.Fatal(err)
	}
}

func runServer() error {
	address, err1 := config.GetServerConnectionData()
	if err1 != nil {
		return err1
	}

	authServer := service.AuthServer{}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	serverRegistrar := grpc.NewServer(opts...)

	api.RegisterAuthServiceServer(serverRegistrar, &authServer)

	api.RegisterDEMServer(serverRegistrar, &service.MyDEMServer{})

	if err = serverRegistrar.Serve(listener); err != nil {
		fmt.Println("failed to serve: %s" + err.Error())
	}

	return nil
}
