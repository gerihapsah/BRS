package main

import (
	"config"
	"db"
	"fmt"
	"log"
	"net/http"
	"os"
	"route"
)

var jconfig = &config.Configuration{}

func main() {
	// Load the configuration file
	config.Load("config"+string(os.PathSeparator)+"config.json", jconfig)
	port := jconfig.Port
	log.Println("Application starting...")
	//connect database
	db.Conect(jconfig.Database)

	log.Print("Server running on port ", port, "...\n")
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", port), route.Routes()))
}
