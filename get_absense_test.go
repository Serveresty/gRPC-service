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
					Id:          17,
					PersonId:    1,
					ReasonId:    11,
					CreatedDate: &timestamppb.Timestamp{Seconds: 1691971200},
					DateFrom:    &timestamppb.Timestamp{Seconds: 1691798400},
					DateTo:      &timestamppb.Timestamp{Seconds: 1691884799},
				},
				{
					Id:          19,
					PersonId:    2,
					ReasonId:    1,
					CreatedDate: &timestamppb.Timestamp{Seconds: 1691971200},
					DateFrom:    &timestamppb.Timestamp{Seconds: 1691798400},
					DateTo:      &timestamppb.Timestamp{Seconds: 1691884799},
				}},
		},
		{
			NameTest: "Get absense by time",
			DateFrom: &timestamppb.Timestamp{Seconds: 1691798400},
			DateTo:   &timestamppb.Timestamp{Seconds: 1691884799},
			Result: []*api.OutputAbsenceData{
				{
					Id:          17,
					PersonId:    1,
					ReasonId:    11,
					CreatedDate: &timestamppb.Timestamp{Seconds: 1691971200},
					DateFrom:    &timestamppb.Timestamp{Seconds: 1691798400},
					DateTo:      &timestamppb.Timestamp{Seconds: 1691884799},
				},
				{
					Id:          19,
					PersonId:    2,
					ReasonId:    1,
					CreatedDate: &timestamppb.Timestamp{Seconds: 1691971200},
					DateFrom:    &timestamppb.Timestamp{Seconds: 1691798400},
					DateTo:      &timestamppb.Timestamp{Seconds: 1691884799},
				},
				{
					Id:          21,
					PersonId:    3,
					ReasonId:    1,
					CreatedDate: &timestamppb.Timestamp{Seconds: 1691971200},
					DateFrom:    &timestamppb.Timestamp{Seconds: 1691798400},
					DateTo:      &timestamppb.Timestamp{Seconds: 1691884799},
				}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.NameTest, func(t *testing.T) {
			req := &api.AbsenceStatusRequest{InputAbsenceData: &api.InputAbsenceData{
				PersonIds: tt.PersonIds,
				DateFrom:  tt.DateFrom,
				DateTo:    tt.DateTo,
			}}
			res, err := c.CheckAbsenceStatus(newCtx, req)
			if err != nil {
				t.Errorf("\nCheckAbsenseTest(%v) got an error: %v", tt.NameTest, err)
			}
			if !reflect.DeepEqual(res.AbsenceData, tt.Result) {
				t.Errorf("\nCheckAbsenseTest(%v)=%v, wanted %v", tt.NameTest, res.AbsenceData, tt.Result)
			}
			t.Logf("\nResult: %v\n", res.AbsenceData)
		})
	}
}
