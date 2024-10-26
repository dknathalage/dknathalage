package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
