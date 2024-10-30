package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// App encapsulates application dependencies.
// This is useful for injecting dependencies in tests.
type App struct{}

// HelloHandler responds with a greeting message.
func (a *App) HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker!")
}

// NewServer configures and returns an http.Server.
func NewServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 3 * time.Second,
	}
}

func main() {
	app := &App{}

	// Create a new ServeMux and register handlers.
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.HelloHandler)

	// Define the server address.
	serverAddr := ":80"
	fmt.Printf("Starting server on %s\n", serverAddr)

	// Initialize the server.
	server := NewServer(serverAddr, mux)

	// Start the server.
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v", err)
	}
}
