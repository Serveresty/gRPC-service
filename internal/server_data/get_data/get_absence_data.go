package getdata

import (
	"encoding/json"
	"os"
	"proteitestcase/pkg/api"
)

func GetAllAbsence() ([]*api.OutputAbsenceData, error) {
	data, err := GetAbsenceData()
	if err != nil {
		return []*api.OutputAbsenceData{}, err
	}

	var absenceData []*api.OutputAbsenceData

	err = json.Unmarshal(data, &absenceData)
	if err != nil {
		return []*api.OutputAbsenceData{}, err
	}

	return absenceData, nil
}

func GetAbsenceByFilter(data *api.InputAbsenceData) ([]*api.OutputAbsenceData, error) {
	absenceData, err := GetAllAbsence()
	if err != nil {
		return []*api.OutputAbsenceData{}, err
	}

	var resultData []*api.OutputAbsenceData

	for _, element := range absenceData {
		if data.DateFrom.AsTime().After(element.DateFrom.AsTime()) || data.DateTo.AsTime().Before(element.DateTo.AsTime()) {
			resultData = append(resultData, element)
		}
		for _, elementData := range data.PersonIds {
			if elementData == element.Id {
				resultData = append(resultData, element)
				break
			}
		}
	}

	return resultData, nil
}

func GetAbsenceData() ([]byte, error) {
	data, err := os.ReadFile("./internal/server_data/absence_data.json")
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}
