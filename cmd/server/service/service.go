package service

import (
	"context"
	"errors"
	"proteitestcase/internal/config"
	getdata "proteitestcase/internal/server_data/get_data"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MyDEMServer struct {
	api.UnimplementedDEMServer
}

func (s *MyDEMServer) Connection(ctx context.Context, req *api.ConnectionRequest) (*api.ConnectionResponse, error) {
	login, password, err := config.GetAuthData()
	if err != nil {
		return &api.ConnectionResponse{
			IsAccessGranted: false,
		}, err
	}

	if (login != req.Login) || (password != req.Password) {
		return &api.ConnectionResponse{
			IsAccessGranted: false,
		}, errors.New("Bad credentials")
	}

	return &api.ConnectionResponse{
		IsAccessGranted: true,
	}, nil
}

func (s *MyDEMServer) GetInfoAboutUser(ctx context.Context, req *api.GetInfoRequest) (*api.GetInfoResponse, error) {
	if (req.UsersData) == (&api.InputUsersData{}) {
		usersData, err := getdata.GetAllUsers()
		if err != nil {
			return &api.GetInfoResponse{
				Status:    status.New(codes.NotFound, "").String(),
				UsersData: []*api.OutputUsersData{},
			}, err
		}
		return &api.GetInfoResponse{
			Status:    status.New(codes.OK, "").String(),
			UsersData: usersData,
		}, nil
	}

	usersData, err := getdata.GetUsersByFilter(req.UsersData)
	if err != nil {
		return &api.GetInfoResponse{
			Status:    status.New(codes.NotFound, "").String(),
			UsersData: []*api.OutputUsersData{},
		}, err
	}

	return &api.GetInfoResponse{
		Status:    status.New(codes.OK, "").String(),
		UsersData: usersData,
	}, nil
}

func (s *MyDEMServer) CheckAbsenceStatus(ctx context.Context, req *api.AbsenceStatusRequest) (*api.AbsenceStatusResponse, error) {
	if (req.InputAbsenceData) == (&api.InputAbsenceData{}) {
		usersData, err := getdata.GetAllAbsence() //Другое
		if err != nil {
			return &api.AbsenceStatusResponse{
				Status:      status.New(codes.NotFound, "").String(),
				AbsenceData: []*api.OutputAbsenceData{},
			}, err
		}
		return &api.AbsenceStatusResponse{
			Status:      status.New(codes.OK, "").String(),
			AbsenceData: usersData, //Из другого
		}, nil
	}

	usersData, err := getdata.GetAbsenceByFilter(req.InputAbsenceData) //Другое
	if err != nil {
		return &api.AbsenceStatusResponse{
			Status:      status.New(codes.NotFound, "").String(),
			AbsenceData: []*api.OutputAbsenceData{},
		}, err
	}

	return &api.AbsenceStatusResponse{
		Status:      status.New(codes.OK, "").String(),
		AbsenceData: usersData, //Из другого
	}, nil
}
