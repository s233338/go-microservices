package main

// Build a webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Create a new router
	router := chi.NewRouter()

	// Log all requests to the console
	router.Use(middleware.Logger)

	// Handle requests to the /hello path
	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

// Handle HTTP routing.
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
