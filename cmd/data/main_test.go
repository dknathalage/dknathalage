package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloHandler verifies that the HelloHandler returns the correct response.
func TestHelloHandler(t *testing.T) {
	app := &App{}

	// Create a new HTTP request.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Create a handler from the HelloHandler method.
	handler := http.HandlerFunc(app.HelloHandler)

	// Serve the HTTP request.
	handler.ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body.
	expected := "Hello, Docker!"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Check the response header.
	if contentType := rr.Header().Get("Content-Type"); contentType != "text/plain; charset=utf-8" {
		t.Errorf("Handler returned unexpected content type: got %v want %v",
			contentType, "text/plain; charset=utf-8")
	}
}

// TestNewServer verifies that the server starts without errors.
// Note: This is a basic test and doesn't fully test server behavior.
func TestNewServer(t *testing.T) {
	// Create a dummy handler.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Initialize the server.
	server := NewServer(":0", handler) // ":0" lets the OS choose an available port.

	// Start the server in a goroutine.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Server failed to start: %v", err)
		}
	}()

	// Give the server a moment to start.
	// In real scenarios, use synchronization mechanisms.
	// Here, we'll use a simple sleep.
	// time.Sleep(100 * time.Millisecond)

	// Since the server is running on ":0", we won't interact with it further.
	// This test ensures that NewServer doesn't panic and starts the server.
}
