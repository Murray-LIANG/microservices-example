package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Murray-LIANG/microservices-example/bookings/common"
	"github.com/Murray-LIANG/microservices-example/bookings/data"
	"github.com/Murray-LIANG/microservices-example/bookings/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// BookingsData defines the bookings list returned by GET /bookings.
type BookingsData struct {
	Data []models.Booking `json:"data"`
}

// BookingData defines the booking used by POST /bookings.
type BookingData struct {
	Data models.Booking `json:"data"`
}

// GetBookings returns the bookings list. It's the handler for HTTP GET of /bookings.
func GetBookings(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext()
	defer ctx.Close()

	mgc := ctx.DBCollection("bookings")
	repo := data.NewBookingRepository(mgc)
	bookings := repo.GetAll()
	j, err := json.Marshal(BookingsData{bookings})
	if err != nil {
		common.ResponseError(
			w, err, "failed to marshal bookings data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// CreateBooking creates the booking. It's the handler for HTTP POST of /bookings.
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking BookingData
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		common.ResponseError(
			w, err, "failed to unmarshal booking data", http.StatusInternalServerError)
		return
	}

	ctx := NewContext()
	defer ctx.Close()

	mgc := ctx.DBCollection("bookings")
	repo := data.NewBookingRepository(mgc)
	if err := repo.Create(&booking.Data); err != nil {
		common.ResponseError(
			w, err, "failed to create booking", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(booking)
	if err != nil {
		common.ResponseError(
			w, err, "failed to marshal booking data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// DeleteBooking deletes the booking by id. It's the handler for HTTP DELETE of
// /bookings/{id}
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	bookingId := requestVars["id"]

	ctx := NewContext()
	defer ctx.Close()

	mgc := ctx.DBCollection("bookings")
	repo := data.NewBookingRepository(mgc)
	if err := repo.Delete(bookingId); err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		common.ResponseError(
			w, err, "failed to delete booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
