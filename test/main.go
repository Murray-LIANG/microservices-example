package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

func main() {

	_, err := mgo.DialWithInfo(
		&mgo.DialInfo{
			Addrs:    []string{"db"},
			Username: "",
			Password: "",
			Timeout:  60 * time.Second,
		},
	)
	if err != nil {
		log.Fatalf("failed to create db session: %s\n", err)
	}
}
