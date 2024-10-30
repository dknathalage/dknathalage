package main

import (
	"log"
	"net/http"

	application "github.com/dknathalage/pkg"
)

func main() {
	app := &application.App{}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.HelloHandler)

	server := application.NewServer(":80", mux)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
