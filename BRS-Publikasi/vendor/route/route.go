package route

import (
	"controller"

	"github.com/gorilla/mux"
)

type Exception struct {
	Message string `json:"message"`
}

func Routes() *mux.Router {

	router := mux.NewRouter()

	//API User
	router.HandleFunc("/api/crawl/{prov}", controller.Crawl).Methods("GET")
	router.HandleFunc("/api/brs", controller.GetBRS).Methods("GET")
	router.HandleFunc("/api/brs/id/{id}", controller.GetBRSById).Methods("GET")
	router.HandleFunc("/api/brs/tanggal/{tanggal}", controller.GetBrsByTanggal).Methods("GET")
	router.HandleFunc("/api/brs/prov/{prov}", controller.GetBrsByProv).Methods("GET")
	router.HandleFunc("/api/brs", controller.PostBRS).Methods("POST")
	return router
}
