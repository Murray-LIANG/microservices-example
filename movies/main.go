package main

import (
	"log"
	"net/http"

	"github.com/Murray-LIANG/microservices-example/movies/common"
	"github.com/Murray-LIANG/microservices-example/movies/routers"
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
