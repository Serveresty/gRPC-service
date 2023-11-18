package logger

import (
	"log"
	"os"
)

func InitHTTP() {
	logFile, err := os.OpenFile("./logs/http_logs", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logFile)
}
