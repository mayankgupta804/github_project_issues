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
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port" + port)
		log.Println("Error: " + err.Error())
	}
}
