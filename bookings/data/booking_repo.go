package data

import (
	"github.com/Murray-LIANG/microservices-example/bookings/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BookingRepository struct {
	mgc *mgo.Collection
}

func NewBookingRepository(mgc *mgo.Collection) *BookingRepository {
	return &BookingRepository{mgc}
}

func (r *BookingRepository) Create(booking *models.Booking) error {
	booking.Id = bson.NewObjectId()
	return r.mgc.Insert(booking)
}

func (r *BookingRepository) GetAll() (bookings []models.Booking) {
	iter := r.mgc.Find(nil).Iter()
	result := models.Booking{}
	for iter.Next(&result) {
		bookings = append(bookings, result)
	}
	return
}

func (r *BookingRepository) Delete(id string) error {
	return r.mgc.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
