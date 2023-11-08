package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"proteitestcase/internal/config/models"
)

func GetConnectionParams(name string) {
	data, err := os.ReadFile("./internal/config/cfg.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	connectionArr := models.ConnectionArr{}

	err = json.Unmarshal(data, &connectionArr)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func whatsDataFromCFG(name string) {

}
