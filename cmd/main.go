package main

import (
	"JWT-tutorial/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Server...")
	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/home", handlers.HomeHandler).Methods(http.MethodGet)

	http.ListenAndServe("127.0.0.1:8080", r)
	fmt.Println("Stopping Server...")
}
