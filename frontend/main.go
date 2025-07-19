package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)
func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	
	// Handle main page
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, nil)
    })
    log.Println("Server starting on http://localhost:8080")
    
	log.Fatal(http.ListenAndServe(":8080", r))
}
