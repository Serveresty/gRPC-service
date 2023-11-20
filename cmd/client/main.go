package main

import (
	"context"
	"proteitestcase/cmd/client/client"
	"proteitestcase/cmd/server/service"
	"proteitestcase/internal/config"
	"proteitestcase/logger"
	"proteitestcase/pkg/api"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crtFile = "./internal/server_data/openssl/server.crt"
)

func main() {
	// Init loggers
	lg := logger.ErrorWarningLogger()

	// Get client connection data from json config
	address, err := config.GetClientConnectionData()
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot get client connection data")
	}

	// Load keys and certificates for TLS
	creds, err := credentials.NewClientTLSFromFile(crtFile, address)
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot get TLS key pairs for client")
	}

	opts := grpc.WithTransportCredentials(creds)

	// Dial grpc client
	cc1, err1 := grpc.Dial(address, opts, grpc.WithUnaryInterceptor(client.Interceptor))
	if err1 != nil {
		lg.Fatal().Err(err).Msg("cannot start grpc connection by client")
	}
	defer cc1.Close()

	c := api.NewAuthServiceClient(cc1)

	login, password, err := config.GetAuthData()
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot get client auth data")
	}

	loginRep, err := c.Login(context.Background(), &api.LoginRequest{Login: login, Password: password})
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot to auth by client")
	}

	requestToken := new(service.AuthToken)
	requestToken.Token = loginRep.Token

	cc2, err := grpc.Dial(address, opts, grpc.WithPerRPCCredentials(requestToken), grpc.WithUnaryInterceptor(client.Interceptor))
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot start grpc connection by client with RPC Credentials")
	}
	defer cc2.Close()

	cl := api.NewDEMClient(cc2)

	info, err := cl.GetInfoAboutUser(context.Background(), &api.GetInfoRequest{UsersData: &api.InputUsersData{}})
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot get info about user")
	}
	log.Print(info)

	abs, err := cl.CheckAbsenceStatus(context.Background(), &api.AbsenceStatusRequest{InputAbsenceData: &api.InputAbsenceData{}})
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot check absence status")
	}
	log.Print(abs)
}
