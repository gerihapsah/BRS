package main

import (
	"config"
	"crawl"
	"db"
	"fmt"
	"log"
	"model"
	"os"
)

var jconfig = &config.Configuration{}

func main() {
	// Load the configuration file
	config.Load("config"+string(os.PathSeparator)+"config.json", jconfig)
	//	port := jconfig.Port
	log.Println("Application starting...")
	//connect database
	db.Conect(jconfig.Database)
	// brs := BRSIndex("http://aceh.bps.go.id")
	brs := crawl.BRSIndex("http://aceh.bps.go.id")
	for _, el := range brs {
		//fmt.Println(el)
		fmt.Println(model.BrsCreate(model.Brs{Judul: el.Judul, Tanggal: el.Tanggal, Abstraksi: el.Abstraksi, Link: el.Link, Prov: "aceh", Website: "http://aceh.bp.go.id"}))
	}
}
