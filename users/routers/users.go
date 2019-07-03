package routers

import (
	"github.com/Murray-LIANG/microservices-example/users/controllers"
	"github.com/gorilla/mux"
)

// SetUsersRouters sets up the routers of users URLs.
func SetUsersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	return router
}
