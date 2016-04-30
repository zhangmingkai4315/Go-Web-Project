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

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

// 统一的错误信息返回
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("[AppError]:%s", handlerError)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}
