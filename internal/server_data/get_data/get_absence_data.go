package getdata

import (
	"encoding/json"
	"os"
	"proteitestcase/internal/server_data/get_data/models"
	"proteitestcase/pkg/api"
)

func GetAllAbsence() ([]*api.OutputAbsenceData, error) {
	data, err := GetAbsenceData()
	if err != nil {
		return []*api.OutputAbsenceData{}, err
	}

	var absData models.GotAbsenceData

	err = json.Unmarshal(data, &absData)
	if err != nil {
		return []*api.OutputAbsenceData{}, err
	}

	return absData.AbsenceData, nil
}

func GetAbsenceByFilter(data *api.InputAbsenceData) ([]*api.OutputAbsenceData, error) {
	absenceData, err := GetAllAbsence()
	if err != nil {
		return []*api.OutputAbsenceData{}, err
	}

	var absData models.GotAbsenceData

	for _, element := range absenceData {
		if data.DateFrom.AsTime().After(element.DateFrom.AsTime()) || data.DateTo.AsTime().Before(element.DateTo.AsTime()) {
			absData.AbsenceData = append(absData.AbsenceData, element)
		}
		for _, elementData := range data.PersonIds {
			if elementData == element.Id {
				absData.AbsenceData = append(absData.AbsenceData, element)
				break
			}
		}
	}

	return absData.AbsenceData, nil
}

func GetAbsenceData() ([]byte, error) {
	data, err := os.ReadFile("./internal/server_data/absence_data.json")
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}
