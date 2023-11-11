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
