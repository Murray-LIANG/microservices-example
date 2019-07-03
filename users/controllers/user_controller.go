package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Murray-LIANG/microservices-example/users/common"
	"github.com/Murray-LIANG/microservices-example/users/data"
	"github.com/Murray-LIANG/microservices-example/users/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// UsersData defines the users list returned by GET /users.
type UsersData struct {
	Data []models.User `json:"data"`
}

// UserData defines the user used by POST /users.
type UserData struct {
	Data models.User `json:"data"`
}

// GetUsers returns the users list. It's the handler for HTTP GET of /users.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext()
	defer ctx.Close()

	mgc := ctx.DBCollection("users")
	repo := data.NewUserRepository(mgc)
	users := repo.GetAll()
	j, err := json.Marshal(UsersData{users})
	if err != nil {
		common.ResponseError(
			w, err, "failed to marshal users data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// CreateUser creates the user. It's the handler for HTTP POST of /users.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserData
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		common.ResponseError(
			w, err, "failed to unmarshal user data", http.StatusInternalServerError)
		return
	}

	ctx := NewContext()
	defer ctx.Close()

	mgc := ctx.DBCollection("users")
	repo := data.NewUserRepository(mgc)
	if err := repo.Create(&user.Data); err != nil {
		common.ResponseError(
			w, err, "failed to create user", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		common.ResponseError(
			w, err, "failed to marshal user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// DeleteUser deletes the user by id. It's the handler for HTTP DELETE of /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	userId := requestVars["id"]

	ctx := NewContext()
	defer ctx.Close()

	mgc := ctx.DBCollection("users")
	repo := data.NewUserRepository(mgc)
	if err := repo.Delete(userId); err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		common.ResponseError(
			w, err, "failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
