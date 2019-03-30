package service

import (
	"log"
	"net/http"
)

// StartWebServer starts the webserver at a given port
func StartWebServer(port string) {
	router := NewRouter()
	http.Handle("/", router)
	log.Println("Starting HTTP service at port:", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("An error occurred starting HTTP listener at port" + port)
		log.Println("Error: " + err.Error())
	}
}
