package tsts

import (
	"context"
	"log"
	"net"
	"proteitestcase/cmd/server/service"
	"proteitestcase/pkg/api"
	"proteitestcase/utils"
	"reflect"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetUsers(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	server := grpc.NewServer()
	t.Cleanup(func() {
		server.Stop()
	})

	demServer := &service.MyDEMServer{}
	api.RegisterDEMServer(server, demServer)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("server.Serve %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	requestToken := new(service.AuthToken)
	requestToken.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6Imhpcm8ifQ.PPSb0vo-Y-OG2teShWxgI0bICn3MthALuvHpPbUvj3E"
	ctx := context.Background()
	md := metadata.New(map[string]string{"authorization": requestToken.Token})
	newCtx := metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.DialContext(newCtx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.DialContext %v", err)
	}
	defer conn.Close()

	c := api.NewDEMClient(conn)

	tests := []struct {
		NameTest  string
		Id        []int64
		Name      string
		WorkPhone string
		Email     string
		DateFrom  *timestamppb.Timestamp
		DateTo    *timestamppb.Timestamp
		Result    []*api.OutputUsersData
	}{
		{
			NameTest: "Get several users by ids",
			Id:       []int64{1, 2},
			Result: []*api.OutputUsersData{
				{
					Id:          1,
					DisplayName: "Иванов Семен Петрович",
					Email:       "petrovich@mail.ru",
					WorkPhone:   "1111"},
				{
					Id:          2,
					DisplayName: "Семенов Петр Иванович",
					Email:       "ivanovich@mail.ru",
					WorkPhone:   "2222"}},
		},
		{
			NameTest: "Get user by email/Check emoji",
			Email:    "petrovich@mail.ru",
			Result: []*api.OutputUsersData{
				{
					Id:          1,
					DisplayName: "Иванов Семен Петрович" + utils.GetEmojiById(11),
					Email:       "petrovich@mail.ru",
					WorkPhone:   "1111"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.NameTest, func(t *testing.T) {
			req := &api.GetInfoRequest{UsersData: &api.InputUsersData{
				Id:        tt.Id,
				Name:      tt.Name,
				WorkPhone: tt.WorkPhone,
				Email:     tt.Email,
				DateFrom:  tt.DateFrom,
				DateTo:    tt.DateTo}}
			res, err := c.GetInfoAboutUser(newCtx, req)
			if err != nil {
				t.Errorf("GetUserTest(%v) got an error: %v", tt.NameTest, err)
			}
			if !reflect.DeepEqual(res.UsersData, tt.Result) {
				t.Errorf("GetUserTest(%v)=%v, wanted %v", tt.NameTest, res.UsersData, tt.Result)
			}
			t.Log(res)
		})
	}

}
