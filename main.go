package main

import (
	"fmt"
	"log"
	"net/http"

	cloudrunmetadatabox "github.com/sinmetalcraft/gcpbox/metadata/cloudrun"
)

func handler(w http.ResponseWriter, r *http.Request) {
	service, err := cloudrunmetadatabox.Service()
	if err != nil {
		log.Println(err)
	}

	revision, err := cloudrunmetadatabox.Revision()
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "Hello! %s.%s", service, revision)
}

func main() {
	log.Println("Start Hello World Server")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
