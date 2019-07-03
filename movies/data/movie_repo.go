package data

import (
	"github.com/Murray-LIANG/microservices-example/movies/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MovieRepository struct {
	mgc *mgo.Collection
}

func NewMovieRepository(mgc *mgo.Collection) *MovieRepository {
	return &MovieRepository{mgc}
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	movie.Id = bson.NewObjectId()
	return r.mgc.Insert(movie)
}

func (r *MovieRepository) GetAll() (movies []models.Movie) {
	iter := r.mgc.Find(nil).Iter()
	result := models.Movie{}
	for iter.Next(&result) {
		movies = append(movies, result)
	}
	return
}

func (r *MovieRepository) Delete(id string) error {
	return r.mgc.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
