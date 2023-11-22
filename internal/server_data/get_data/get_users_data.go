package getdata

import (
	"encoding/json"
	"os"
	"proteitestcase/internal/server_data/get_data/models"
	"proteitestcase/pkg/api"
	"proteitestcase/utils"
)

func GetAllUsers() ([]*api.OutputUsersData, error) {
	data, err := GetUsersData()
	if err != nil {
		return []*api.OutputUsersData{}, err
	}

	var dt models.GotUsersData

	err = json.Unmarshal(data, &dt)
	if err != nil {
		return []*api.OutputUsersData{}, err
	}

	return dt.UsersData, nil
}

func GetUsersByFilter(data *api.InputUsersData) ([]*api.OutputUsersData, error) {
	usersData, err := GetAllUsers()
	if err != nil {
		return []*api.OutputUsersData{}, err
	}

	var dt models.GotUsersData

	for _, element := range usersData {
		if (data.Name == element.DisplayName) || (data.WorkPhone == element.WorkPhone) || (data.Email == element.Email) {
			if data.Email != "" {
				var reason int64
				abs, err := GetAbsenceByFilter(&api.InputAbsenceData{PersonIds: []int64{element.Id}})
				if err != nil {
					return []*api.OutputUsersData{}, err
				}
				for _, k := range abs {
					reason = k.ReasonId
				}

				element.DisplayName += utils.GetEmojiById(reason)
			}
			dt.UsersData = append(dt.UsersData, element)
		}
		for _, elementData := range data.Id {
			if elementData == element.Id {
				dt.UsersData = append(dt.UsersData, element)
				break
			}
		}
	}

	return dt.UsersData, nil
}

func GetUsersData() ([]byte, error) {
	data, err := os.ReadFile("./internal/server_data/users_data.json")
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}
