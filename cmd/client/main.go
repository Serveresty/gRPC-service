package main

import (
	"log"
	"proteitestcase/internal/config"
)

func main() {
	if err := runClient(); err != nil {
		log.Fatal(err)
	}
}

func runClient() error {
	address, err1 := config.GetClientConnectionData()
	if err1 != nil {
		return err1
	}
	if address != "" {
		return nil
	}
	return nil
}
