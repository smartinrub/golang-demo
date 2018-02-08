package main

import (
	"html/template"
	"log"
	"net/http"
)

// Handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

// Main
func main() {
	// Routers
	http.HandleFunc("/", indexHandler)

	// open port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
