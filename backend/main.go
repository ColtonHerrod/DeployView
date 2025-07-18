package main

import (
	"deployview-backend/handlers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger" // http-swagger middleware
	_ "deployview-backend/docs"
)

// @title DeployView API
// @version 1.0
// @description This is the API documentation for DeployView, a deployment management tool.
// @host localhost:8080
// @BasePath /
// @schemes http
// @contact.name API Support
// @contact.email
func main() {
	r := mux.NewRouter()

    // Routes consist of a path and a handler function.
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/documentation/", http.StatusMovedPermanently)
	})
	r.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/deployments", handlers.AllDeploymentsHandler).Methods("GET")
	r.HandleFunc("/deployments/{account}", handlers.AccountDeploymentsHandler).Methods("GET")
	r.HandleFunc("/deployments/{account}", handlers.CreateDeploymentHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
