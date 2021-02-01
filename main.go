package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	log.Println("Start Hello World Server")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
