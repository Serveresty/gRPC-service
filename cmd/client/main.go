package main

import (
	"context"
	"fmt"
	"log"
	"proteitestcase/cmd/server/service"
	"proteitestcase/internal/config"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crtFile = "./internal/server_data/openssl/server.crt"
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

	creds, err := credentials.NewClientTLSFromFile(crtFile, address)
	if err != nil {
		return err
	}

	opts := grpc.WithTransportCredentials(creds)

	cc1, err1 := grpc.Dial(address, opts)
	if err1 != nil {
		return err1
	}
	defer cc1.Close()

	c := api.NewAuthServiceClient(cc1)

	login, password, err := config.GetAuthData()
	if err != nil {
		return err
	}

	loginRep, err := c.Login(context.Background(), &api.LoginRequest{Login: login, Password: password})
	if err != nil {
		return err
	}

	requestToken := new(service.AuthToken)
	requestToken.Token = loginRep.Token

	cc2, err := grpc.Dial(address, opts, grpc.WithPerRPCCredentials(requestToken))
	if err != nil {
		return err
	}
	defer cc2.Close()

	cl := api.NewDEMClient(cc2)

	result, err := cl.GetInfoAboutUser(context.Background(), &api.GetInfoRequest{UsersData: &api.InputUsersData{}})
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}
