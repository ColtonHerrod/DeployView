package main

import (
	"deployview-backend/handlers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

    // Routes consist of a path and a handler function.
    r.HandleFunc("/", handlers.RootHandler).Methods("GET")
	r.HandleFunc("/deployments", handlers.AllDeploymentsHandler).Methods("GET")
	r.HandleFunc("/deployments/{account}", handlers.DeploymentsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
