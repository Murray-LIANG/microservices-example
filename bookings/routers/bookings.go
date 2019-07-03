package routers

import (
	"github.com/Murray-LIANG/microservices-example/bookings/controllers"
	"github.com/gorilla/mux"
)

// SetBookingsRouters sets up the routers of bookings URLs.
func SetBookingsRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/bookings", controllers.GetBookings).Methods("GET")
	router.HandleFunc("/bookings", controllers.CreateBooking).Methods("POST")
	router.HandleFunc("/bookings/{id}", controllers.DeleteBooking).Methods("DELETE")
	return router
}
