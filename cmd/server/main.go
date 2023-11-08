package main

import (
	"context"
	"log"
	"net"
	"proteitestcase/internal/config"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc"
)

type myDEMServer struct {
	api.UnimplementedDEMServer
}

func Connection(context.Context, *api.ConnectionRequest) (*api.ConnectionResponse, error) {
	return &api.ConnectionResponse{
		IsAccessGranted: true,
	}, nil
}

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
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	serverRegistrar := grpc.NewServer()
	service := &myDEMServer{}
	api.RegisterDEMServer(serverRegistrar, service)
	err = serverRegistrar.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
