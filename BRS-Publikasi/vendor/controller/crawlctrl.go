package controller

import (
	"encoding/json"
	// "log"
	"crawl"
	"model"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/gorilla/mux"
)

func Crawl(w http.ResponseWriter, r *http.Request) {
	// user, err := model.UserAll()
	// if err != nil {
	// 	log.Println(err)
	// 	Error500(w, r)
	// }
	params := mux.Vars(r)
	prov := params["prov"]
	url := "https://" + prov + ".bps.go.id"
	brs := crawl.BRSIndex(url)
	for _, el := range brs {

		_ = model.BrsCreate(model.Brs{Judul: el.Judul, Tanggal: el.Tanggal, Abstraksi: el.Abstraksi, Link: el.Link, Prov: prov, Website: url})
	}

	json.NewEncoder(w).Encode(brs)
}
