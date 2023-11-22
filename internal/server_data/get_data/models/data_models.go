package models

import (
	"proteitestcase/pkg/api"
)

type GotUsersData struct {
	UsersData []*api.OutputUsersData `json:"usersData"`
}

type GotAbsenceData struct {
	AbsenceData []*api.OutputAbsenceData `json:"absenceData"`
}

type GAbsData struct {
	AbsenceData []AbsenceData `json:"absenceData"`
}

type AbsenceData struct {
	CreatedDate string `json:"createdDate"`
	DateFrom    string `json:"dateFrom"`
	DateTo      string `json:"dateTo"`
	Id          int64  `json:"id"`
	PersonId    int64  `json:"personId"`
	ReasonId    int64  `json:"reasonId"`
}
