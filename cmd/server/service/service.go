package service

import (
	"context"
	"fmt"
	getdata "proteitestcase/internal/server_data/get_data"
	"proteitestcase/logger"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MyDEMServer struct {
	api.UnimplementedDEMServer
}

var (
	lg      = logger.ErrorWarningLogger()
	debugLg = logger.DebugLogger()
)

func (s *MyDEMServer) GetInfoAboutUser(ctx context.Context, req *api.GetInfoRequest) (*api.GetInfoResponse, error) {
	_, err := CheckAuth(ctx)
	if err != nil {
		lg.Err(err).Msg("Not authenticated")
		return nil, fmt.Errorf("Not authenticated: %v", err)
	}
	debugLg.Debug().Str("action", "Checked auth in method 'GetInfoAboutUser'")

	if (req.UsersData.Id == nil) && (req.UsersData.Name == "") && (req.UsersData.Email == "") && (req.UsersData.WorkPhone == "") {
		usersData, err1 := getdata.GetAllUsers()
		if err1 != nil {
			lg.Err(err).Msg("Error while getting all users from json")
			return &api.GetInfoResponse{
				Status:    status.New(codes.NotFound, "").String(),
				UsersData: []*api.OutputUsersData{},
			}, err1
		}

		debugLg.Debug().Str("action", "Got all users from json")
		return &api.GetInfoResponse{
			Status:    status.New(codes.OK, "").String(),
			UsersData: usersData,
		}, nil
	}

	usersData, err := getdata.GetUsersByFilter(req.UsersData)
	if err != nil {
		lg.Err(err).Msg("Error while getting users by filter from json")
		return &api.GetInfoResponse{
			Status:    status.New(codes.NotFound, "").String(),
			UsersData: []*api.OutputUsersData{},
		}, err
	}
	debugLg.Debug().Str("action", "Got users by filter")

	return &api.GetInfoResponse{
		Status:    status.New(codes.OK, "").String(),
		UsersData: usersData,
	}, nil
}

func (s *MyDEMServer) CheckAbsenceStatus(ctx context.Context, req *api.AbsenceStatusRequest) (*api.AbsenceStatusResponse, error) {
	_, err := CheckAuth(ctx)
	if err != nil {
		lg.Err(err).Msg("Not authenticated")
		return nil, fmt.Errorf("Not authenticated: %v", err)
	}
	debugLg.Debug().Str("action", "Checked auth in method 'CheckAbsenceStatus'")

	if (req.InputAbsenceData.DateFrom == nil) && (req.InputAbsenceData.PersonIds == nil) && (req.InputAbsenceData.DateTo == nil) {
		usersData, err1 := getdata.GetAllAbsence()
		if err1 != nil {
			lg.Err(err).Msg("Error while getting all absenses")
			return &api.AbsenceStatusResponse{
				Status:      status.New(codes.NotFound, "").String(),
				AbsenceData: []*api.OutputAbsenceData{},
			}, err1
		}
		debugLg.Debug().Str("action", "Got all absence from json")

		return &api.AbsenceStatusResponse{
			Status:      status.New(codes.OK, "").String(),
			AbsenceData: usersData,
		}, nil
	}

	usersData, err := getdata.GetAbsenceByFilter(req.InputAbsenceData)
	if err != nil {
		lg.Err(err).Msg("Error while getting absenses by filter")
		return &api.AbsenceStatusResponse{
			Status:      status.New(codes.NotFound, "").String(),
			AbsenceData: []*api.OutputAbsenceData{},
		}, err
	}
	debugLg.Debug().Str("action", "Got absenses by filter from json")

	return &api.AbsenceStatusResponse{
		Status:      status.New(codes.OK, "").String(),
		AbsenceData: usersData,
	}, nil
}
