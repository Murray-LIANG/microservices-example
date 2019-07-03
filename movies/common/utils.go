package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

type configuration struct {
	Server      string
	MongoHost   string
	MongoUser   string
	MongoPasswd string
	Database    string
}

// AppConfig holds the configuration values from config.json file.
var AppConfig configuration

func initConfig() {
	f, err := os.Open("common/config.json")
	if err != nil {
		log.Fatalf("failed to open config file: %s\n", err)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("failed to load config: %s\n", err)
	}
}

var dbSession *mgo.Session

func createDBSession() {
	var err error
	dbSession, err = mgo.DialWithInfo(
		&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoHost},
			Username: AppConfig.MongoUser,
			Password: AppConfig.MongoPasswd,
			Timeout:  60 * time.Second,
		},
	)
	if err != nil {
		log.Fatalf("failed to create db session: %s\n", err)
	}
}

func GetDBSession() *mgo.Session {
	if dbSession == nil {
		createDBSession()
	}
	return dbSession
}

type ErrorData struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	HttpStatus int    `json:"status"`
}

type ErrorResponse struct {
	Data ErrorData `json:"data"`
}

func ResponseError(w http.ResponseWriter, err error, message string, code int) {
	data := ErrorData{
		Error:      err.Error(),
		Message:    message,
		HttpStatus: code,
	}

	log.Printf("error response: %s\n", err)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(ErrorResponse{Data: data}); err == nil {
		w.Write(j)
	}
}
