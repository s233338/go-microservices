package main

// Build a webserver

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

// Handle HTTP routing.
func basicHandler(w http.ResponseWriter, r *http.Request) {

}
