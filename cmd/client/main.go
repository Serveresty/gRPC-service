package main

import (
	"context"
	"fmt"
	"log"
	"proteitestcase/internal/config"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc"
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

	conn, err1 := grpc.Dial(address, grpc.WithInsecure())
	if err1 != nil {
		return err1
	}
	defer conn.Close()

	c := api.NewDEMClient(conn)

	login, password, err2 := config.GetAuthData()
	if err2 != nil {
		return err2
	}

	res, err3 := c.Connection(context.Background(), &api.ConnectionRequest{Login: login, Password: password})
	if err3 != nil {
		return err3
	}

	fmt.Println("Is access granted: ")
	fmt.Println(res.IsAccessGranted)

	if res.IsAccessGranted {
		rs, er := c.GetInfoAboutUser(context.Background(), &api.GetInfoRequest{UsersData: &api.InputUsersData{}})
		if er != nil {
			return er
		}
		fmt.Println(rs.Status)
		fmt.Println(rs.UsersData)
	}

	rrs, e := c.CheckAbsenceStatus(context.Background(), &api.AbsenceStatusRequest{InputAbsenceData: &api.InputAbsenceData{}})
	if e != nil {
		return e
	}

	fmt.Println(rrs.Status)
	fmt.Println(rrs.AbsenceData)

	return nil
}
