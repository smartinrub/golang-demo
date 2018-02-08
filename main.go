package main

import (
	"html/template"
	"log"
	"net/http"
)

// Handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Read content of html file and returns a Template
	t, _ := template.ParseFiles("index.html")
	// Execute the template, writing the generated HTML to the http.ResponseWriter
	t.Execute(w, nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// Read content of html file and returns a Template
	t, _ := template.ParseFiles("weather.html")
	// Execute the template, writing the generated HTML to the http.ResponseWriter
	t.Execute(w, nil)
}

// Main
func main() {
	// Routers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/weather", weatherHandler)

	// setting listening port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
