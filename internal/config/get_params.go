package config

import (
	"encoding/json"
	"os"
	"proteitestcase/internal/config/models"
)

func GetServerConnectionData() (string, error) {
	connect := models.ServerConnection{}
	data, err := GetCfg()
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(data, &connect)
	if err != nil {
		return "", err
	}
	return connect.ConData.IP + ":" + connect.ConData.Port, nil
}

func GetClientConnectionData() (string, error) {
	connect := models.ClientConnection{}
	data, err := GetCfg()
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(data, &connect)
	if err != nil {
		return "", err
	}
	return connect.ConData.IP + ":" + connect.ConData.Port, nil
}

func GetAuthData() (string, string, error) {
	connect := models.AuthArr{}
	data, err := GetCfg()
	if err != nil {
		return "", "", err
	}

	err = json.Unmarshal(data, &connect)
	if err != nil {
		return "", "", err
	}
	return connect.AuData.Login, connect.AuData.Password, nil
}

func GetCfg() ([]byte, error) {
	data, err := os.ReadFile("./internal/config/cfg.json")
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}
