package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

var AppConfig configuration

func initConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loading file :%s\n]", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err := decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loading file :%s\n]", err)
	}
}
