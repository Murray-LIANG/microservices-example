package data

import (
	"github.com/Murray-LIANG/microservices-example/users/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	mgc *mgo.Collection
}

func NewUserRepository(mgc *mgo.Collection) *UserRepository {
	return &UserRepository{mgc}
}

func (r *UserRepository) Create(user *models.User) error {
	user.Id = bson.NewObjectId()
	return r.mgc.Insert(user)
}

func (r *UserRepository) GetAll() (users []models.User) {
	iter := r.mgc.Find(nil).Iter()
	result := models.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return
}

func (r *UserRepository) Delete(id string) error {
	return r.mgc.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
