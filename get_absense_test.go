package tsts

import (
	"context"
	"log"
	"net"
	"proteitestcase/cmd/server/service"
	"proteitestcase/pkg/api"
	"reflect"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestAbsenseStatus(t *testing.T) {
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
		PersonIds []int64
		DateFrom  *timestamppb.Timestamp
		DateTo    *timestamppb.Timestamp
		Result    []*api.OutputAbsenceData
	}{
		{
			NameTest:  "Get several absenses by person's ids",
			PersonIds: []int64{1, 2},
			Result: []*api.OutputAbsenceData{
				{
					Id:       17,
					PersonId: 1,
					ReasonId: 11,
				},
				{
					Id:       19,
					PersonId: 2,
					ReasonId: 1,
				}},
		},
		/* {
			NameTest: "Get user by email/Check emoji",
			Result:   []*api.OutputAbsenceData{},
		}, */
	}

	for _, tt := range tests {
		t.Run(tt.NameTest, func(t *testing.T) {
			req := &api.AbsenceStatusRequest{InputAbsenceData: &api.InputAbsenceData{
				PersonIds: tt.PersonIds,
			}}
			res, err := c.CheckAbsenceStatus(newCtx, req)
			if err != nil {
				t.Errorf("CheckAbsenseTest(%v) got an error: %v", tt.NameTest, err)
			}
			if !reflect.DeepEqual(res.AbsenceData, tt.Result) {
				t.Errorf("CheckAbsenseTest(%v)=%v, wanted %v", tt.NameTest, res.AbsenceData, tt.Result)
			}
			t.Log(res)
		})
	}
}
