package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Murray-LIANG/microservices-example/movies/common"
	"github.com/Murray-LIANG/microservices-example/movies/data"
	"github.com/Murray-LIANG/microservices-example/movies/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// MoviesData defines the movies list returned by GET /movies.
type MoviesData struct {
	Data []models.Movie `json:"data"`
}

// MovieData defines the movie used by POST /movies.
type MovieData struct {
	Data models.Movie `json:"data"`
}

// GetMovies returns the movies list. It's the handler for HTTP GET of /movies.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext()
	defer ctx.Close()

	mgc := ctx.DBCollection("movies")
	repo := data.NewMovieRepository(mgc)
	movies := repo.GetAll()
	j, err := json.Marshal(MoviesData{movies})
	if err != nil {
		common.ResponseError(
			w, err, "failed to marshal movies data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// CreateMovie creates the movie. It's the handler for HTTP POST of /movies.
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie MovieData
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		common.ResponseError(
			w, err, "failed to unmarshal movie data", http.StatusInternalServerError)
		return
	}

	ctx := NewContext()
	defer ctx.Close()

	mgc := ctx.DBCollection("movies")
	repo := data.NewMovieRepository(mgc)
	if err := repo.Create(&movie.Data); err != nil {
		common.ResponseError(
			w, err, "failed to create movie", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(movie)
	if err != nil {
		common.ResponseError(
			w, err, "failed to marshal movie data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// DeleteMovie deletes the movie by id. It's the handler for HTTP DELETE of /movies/{id}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	movieId := requestVars["id"]

	ctx := NewContext()
	defer ctx.Close()

	mgc := ctx.DBCollection("movies")
	repo := data.NewMovieRepository(mgc)
	if err := repo.Delete(movieId); err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		common.ResponseError(
			w, err, "failed to delete movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
