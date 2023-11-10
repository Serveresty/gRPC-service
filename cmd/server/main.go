package main

import (
	"fmt"
	"log"
	"net"
	"proteitestcase/cmd/server/service"
	"proteitestcase/internal/config"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc"
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

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	serverRegistrar := grpc.NewServer()
	api.RegisterDEMServer(serverRegistrar, &service.MyDEMServer{})

	if err = serverRegistrar.Serve(listener); err != nil {
		fmt.Println("failed to serve: %s" + err.Error())
	}

	return nil
}
