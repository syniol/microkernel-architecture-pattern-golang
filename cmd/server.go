package main

import (
	"log"
	"net/http"

	"architecture/internal"
)

func main() {
	println("Starting Demo Server for Micro")

	server := http.NewServeMux()
	internal.NewRESTfulAPI(server)

	err := http.ListenAndServe(":8000", server)
	if err != nil {
		log.Fatal(err)
	}
}
