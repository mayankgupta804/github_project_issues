package service

import (
	"log"
	"net/http"
)

// StartWebServer starts the webserver at a given port
func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)
	http.Handle("/issues", r)
	log.Println("Starting HTTP service at port:", port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port" + port)
		log.Println("Error: " + err.Error())
	}
}
