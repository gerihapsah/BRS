package controller

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

// Error404 handles 404 - Page Not Found
func Error404(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Message{Code: 404, Msg: "API Not Found"})
}

// Error500 handles 500 - Internal Server Error
func Error500(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Message{Code: 404, Msg: "Internal Server Error 500"})
}
