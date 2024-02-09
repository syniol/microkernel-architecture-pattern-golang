package main

import (
	"log"
	"net/http"

	"architecture/internal"
)

func main() {
	println("Microkernel Architecture Pattern Demo")
	println("Created by: Hadi Tajallaei")
	println("Copyright (c) Syniol Limited. All rights reserved.")

	server := http.NewServeMux()
	internal.NewRESTfulAPI(server)

	println("\n🚀 Access HTTP Server from: 127.0.0.1:8000")
	err := http.ListenAndServe(":8000", server)
	if err != nil {
		log.Fatal(err)
	}

}
