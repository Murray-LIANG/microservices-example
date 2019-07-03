package main

import (
	"log"
	"net/http"

	"github.com/Murray-LIANG/microservices-example/users/common"
	"github.com/Murray-LIANG/microservices-example/users/routers"
)

func main() {
	common.StartUp()

	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	log.Println("Listening ...")
	server.ListenAndServe()
}
