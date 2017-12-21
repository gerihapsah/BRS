package controller

import (
	"encoding/json"
	"log"
	"model"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := model.UserAll()
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	user, err := model.UserById(id)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]

	user, err := model.UserByEmail(email)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// user.Password = passhash.HashString(user.Password)
	err := model.UserCreate(user)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}
func PutUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := model.UserUpdate(id, user)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := model.UserDelete(id)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(Message{Code: 200, Msg: "Deleted..."})
}
