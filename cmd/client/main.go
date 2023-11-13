package main

import (
	"context"
	"fmt"
	"log"
	"proteitestcase/internal/config"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	hostname = "localhost"
	crtFile  = "./internal/server_data/openssl/server.crt"
)

func main() {
	if err := runClient(); err != nil {
		log.Fatal(err)
	}
}

func runClient() error {
	address, err := config.GetClientConnectionData()
	if err != nil {
		return err
	}

	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err1 := grpc.Dial(address, opts...)
	if err1 != nil {
		return err1
	}
	defer conn.Close()

	c2 := api.NewAuthServiceClient(conn)

	login, password, err := config.GetAuthData()
	if err != nil {
		return err
	}

	resp2, err := c2.Login(context.Background(), &api.LoginRequest{Login: login, Password: password})
	if err != nil {
		return err
	}

	fmt.Println(resp2)

	c := api.NewDEMClient(conn)

	resp, err := c.GetInfoAboutUser(context.Background(), &api.GetInfoRequest{UsersData: &api.InputUsersData{}})
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
