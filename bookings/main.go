package main

import (
	"log"
	"net/http"

	"github.com/Murray-LIANG/microservices-example/bookings/common"
	"github.com/Murray-LIANG/microservices-example/bookings/routers"
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
