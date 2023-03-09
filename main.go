package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	for k, v := range r.Header {
		log.Printf("%s:%v\n", k, v)
	}

	responseStatus := r.FormValue("responseStatus")
	if responseStatus != "" {
		sts, err := strconv.ParseInt(responseStatus, 10, 32)
		if err != nil {
			fmt.Fprintf(w, "Hello! %s.%s.err=%s", service, revision, err)
			return
		}
		w.WriteHeader(int(sts))
	}

	fmt.Fprintf(w, "Hello! %s.%s", service, revision)
}

func main() {
	log.Println("Start Hello World Server")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
