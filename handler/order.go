package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

// Http handler interface
func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create order")
}

func (o Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List all orders")
}

func (o Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get order by ID")
}

func (o Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update order by ID")
}

func (o Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete order  by ID")
}
