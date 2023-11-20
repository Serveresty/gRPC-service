package main

import (
	"crypto/tls"
	"net"
	"proteitestcase/cmd/server/service"
	"proteitestcase/internal/config"
	"proteitestcase/logger"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crtFile = "./internal/server_data/openssl/server.crt"
	keyFile = "./internal/server_data/openssl/server.key"
)

func main() {
	// Init loggers
	lg := logger.ErrorWarningLogger()
	debugLg := logger.DebugLogger()

	// Get server connection data from json config
	serverConData, err1 := config.GetServerConnectionData()
	if err1 != nil {
		lg.Fatal().Err(err1).Msg("cannot get server connection data")
	}
	debugLg.Debug().Str("action", "get server connection data from json config success")

	// Listen server
	listener, err := net.Listen("tcp", serverConData.ConData.IP+":"+serverConData.ConData.Port)
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot start gRPC server")
	}
	debugLg.Debug().Msgf("start gRPC server at %s", listener.Addr().String())

	// Load keys and certificates for TLS
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot get TLS key pairs for server")
	}
	debugLg.Debug().Msg("loaded TLS key pairs")

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(service.GRPCLogger),
	}

	// New gRPC server
	serverRegistrar := grpc.NewServer(opts...)
	debugLg.Debug().Msg("start gRPC server")

	// Init DEM server struct
	demServer := &service.MyDEMServer{}

	// Init Auth server struct
	authServer := service.AuthServer{}

	api.RegisterAuthServiceServer(serverRegistrar, &authServer)
	api.RegisterDEMServer(serverRegistrar, demServer)

	if err = serverRegistrar.Serve(listener); err != nil {
		lg.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}
