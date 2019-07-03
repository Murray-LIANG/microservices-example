package routers

import (
	"github.com/Murray-LIANG/microservices-example/movies/controllers"
	"github.com/gorilla/mux"
)

// SetMoviesRouters sets up the routers of movies URLs.
func SetMoviesRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
	return router
}
