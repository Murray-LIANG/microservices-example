package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes initializes users router.
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	return SetMoviesRouters(router)
}
