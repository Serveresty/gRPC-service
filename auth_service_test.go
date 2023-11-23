package tsts

import (
	"context"
	"errors"
	"log"
	"net"
	"proteitestcase/cmd/server/service"
	"proteitestcase/pkg/api"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestLogin(t *testing.T) {

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	server := grpc.NewServer()
	t.Cleanup(func() {
		server.Stop()
	})

	authServer := service.AuthServer{}
	api.RegisterAuthServiceServer(server, &authServer)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("server.Serve %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.DialContext %v", err)
	}
	defer conn.Close()

	c := api.NewAuthServiceClient(conn)

	tests := []struct {
		login    string
		password string
		err      error
	}{
		{
			login:    "hiro",
			password: "qwerty",
			err:      errors.New(""),
		},
		{
			login:    "kuroko",
			password: "zxcvbn",
			err:      errors.New("Bad credentials"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.login, func(t *testing.T) {
			req := &api.LoginRequest{Login: tt.login, Password: tt.password}
			_, err := c.Login(context.Background(), req)
			if errors.Is(err, tt.err) {
				t.Errorf("LoginTest(%v)=%v, wanted %v", tt.login, err, tt.err)
			}
		})
	}
}
