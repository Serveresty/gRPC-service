package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"proteitestcase/cmd/server/service"
	"proteitestcase/internal/config"
	"proteitestcase/internal/logger"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crtFile = "./internal/server_data/openssl/server.crt"
	keyFile = "./internal/server_data/openssl/server.key"
)

func main() {
	logger.InitHTTP()
	if err := runServer(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

func runServer() error {
	serverConData, err1 := config.GetServerConnectionData()
	if err1 != nil {
		return err1
	}

	authServer := service.AuthServer{}

	listener, err := net.Listen("tcp", serverConData.ConData.IP+":"+serverConData.ConData.Port)
	if err != nil {
		return err
	}

	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(service.Interceptor),
	}

	serverRegistrar := grpc.NewServer(opts...)

	demServer := &service.MyDEMServer{}

	api.RegisterAuthServiceServer(serverRegistrar, &authServer)

	api.RegisterDEMServer(serverRegistrar, demServer)

	if err = serverRegistrar.Serve(listener); err != nil {
		fmt.Println("failed to serve: %s" + err.Error())
	}

	return nil
}
