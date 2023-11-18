package main

import (
	"crypto/tls"
	"net"
	"os"
	"proteitestcase/cmd/server/service"
	"proteitestcase/internal/config"
	"proteitestcase/pkg/api"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crtFile = "./internal/server_data/openssl/server.crt"
	keyFile = "./internal/server_data/openssl/server.key"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	serverConData, err1 := config.GetServerConnectionData()
	if err1 != nil {
		log.Fatal().Err(err1).Msg("cannot get server connection data")
	}

	authServer := service.AuthServer{}

	listener, err := net.Listen("tcp", serverConData.ConData.IP+":"+serverConData.ConData.Port)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())

	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get TLS key pairs for server")
	}
	log.Info().Msgf("loaded TLS key pairs")

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(service.GRPCLogger),
	}

	serverRegistrar := grpc.NewServer(opts...)
	log.Info().Msg("start gRPC server")

	demServer := &service.MyDEMServer{}

	api.RegisterAuthServiceServer(serverRegistrar, &authServer)

	api.RegisterDEMServer(serverRegistrar, demServer)

	if err = serverRegistrar.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}
