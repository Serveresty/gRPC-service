package getdata

import (
	"encoding/json"
	"os"
	"proteitestcase/internal/server_data/get_data/models"
	"proteitestcase/pkg/api"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	layout     = "2006-01-02T15:04:05"
	dataLayout = "2006-01-02"
)

func GetAllAbsence() ([]*api.OutputAbsenceData, error) {
	data, err := GetAbsenceData()
	if err != nil {
		return []*api.OutputAbsenceData{}, err
	}

	var absData models.GotAbsenceData

	var aData models.GAbsData

	err = json.Unmarshal(data, &aData)
	if err != nil {
		return []*api.OutputAbsenceData{}, err
	}

	for _, element := range aData.AbsenceData {
		createdDate, err := time.Parse(dataLayout, element.CreatedDate)
		if err != nil {
			continue
		}
		dateFrom, err := time.Parse(layout, element.DateFrom)
		if err != nil {
			continue
		}
		dateTo, err := time.Parse(layout, element.DateTo)
		if err != nil {
			continue
		}
		absData.AbsenceData = append(absData.AbsenceData,
			&api.OutputAbsenceData{
				Id:          element.Id,
				PersonId:    element.PersonId,
				CreatedDate: timestamppb.New(createdDate),
				DateFrom:    timestamppb.New(dateFrom),
				DateTo:      timestamppb.New(dateTo),
			})
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
		if data.DateFrom.AsTime().After(element.DateFrom.AsTime()) && data.DateTo.AsTime().Before(element.DateTo.AsTime()) {
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
