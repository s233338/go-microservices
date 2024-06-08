package main

// Build a webserver

import (
	"context"
	"fmt"

	"github.com/s233338/go-microservices/application"
)

func main() {
	app := application.New()

	err := app.Start(context.Background())
	if err != nil {
		fmt.Printf("failed to start server: %v\n", err)
	}
}
