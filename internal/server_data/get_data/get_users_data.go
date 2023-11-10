package getdata

import (
	"encoding/json"
	"os"
	"proteitestcase/pkg/api"
)

func GetAllUsers() ([]*api.OutputUsersData, error) {
	data, err := GetUsersData()
	if err != nil {
		return []*api.OutputUsersData{}, err
	}

	var usersData []*api.OutputUsersData

	err = json.Unmarshal(data, &usersData)
	if err != nil {
		return []*api.OutputUsersData{}, err
	}

	return usersData, nil
}

func GetUsersByFilter(data *api.InputUsersData) ([]*api.OutputUsersData, error) {
	usersData, err := GetAllUsers()
	if err != nil {
		return []*api.OutputUsersData{}, err
	}

	var resultData []*api.OutputUsersData

	for _, element := range usersData {
		if (data.Name == element.DisplayName) || (data.WorkPhone == element.WorkPhone) || (data.Email == element.Email) {
			resultData = append(resultData, element)
		}
		for _, elementData := range data.Id {
			if elementData == element.Id {
				resultData = append(resultData, element)
				break
			}
		}
	}

	return resultData, nil
}

func GetUsersData() ([]byte, error) {
	data, err := os.ReadFile("./internal/server_data/users_data.json")
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}
