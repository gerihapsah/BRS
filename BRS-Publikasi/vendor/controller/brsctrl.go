package controller

import (
	"encoding/json"
	"log"
	"model"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBRS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := model.BrsAll()
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}

func GetBRSById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	user, err := model.BrsById(id)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}
func GetBrsByTanggal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["tanggal"]
	user, err := model.BrsByTanggal(id)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}
func GetBrsByProv(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["prov"]
	user, err := model.BrsByProv(id)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}

func PostBRS(w http.ResponseWriter, r *http.Request) {
	var user model.Brs
	_ = json.NewDecoder(r.Body).Decode(&user)

	// user.Password = passhash.HashString(user.Password)
	err := model.BrsCreate(user)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(user)
}
func PutBRS(w http.ResponseWriter, r *http.Request) {
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
func DeleteBRS(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := model.UserDelete(id)
	if err != nil {
		log.Println(err)
		Error500(w, r)
	}
	json.NewEncoder(w).Encode(Message{Code: 200, Msg: "Deleted..."})
}
