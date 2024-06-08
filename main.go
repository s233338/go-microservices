package main

// Build a webserver

import (
	"context"
	"fmt"
	"go-microservices/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("failed to start server: %v\n", err)
	}
}
