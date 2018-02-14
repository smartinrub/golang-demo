package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Data Structures
// Weather
type Data struct {
	Temperature float32 `json:"temp"`
	TempMin     float32 `json:"temp_min"`
	TempMax     float32 `json:"temp_max"`
	Humidity    float32 `json:"humidity"`
}

// City Data
type City struct {
	Name     string `json:"name"`
	MainData Data   `json:"main"`
}

// Handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Read content of html file and returns a Template
	t, _ := template.ParseFiles("index.html")
	// Execute the template, writing the generated HTML to the http.ResponseWriter
	t.Execute(w, nil)

	elapsed := time.Since(start)
	log.Printf("Weather page took %s", elapsed)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieves value from form
	cityName := r.FormValue("city")

	// Build url
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?APPID=bd5e378503939ddaee76f12ad7a97608&q=%s&mode=json&units=metric", cityName)

	start := time.Now()

	// Call REST API
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	}

	// Parse data into Struct
	data, _ := ioutil.ReadAll(response.Body)
	textBytes := []byte(data)
	cityValues := City{}
	jsonErr := json.Unmarshal(textBytes, &cityValues)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return
	}

	// If city name not found by API
	if cityValues.Name == "" {
		t, _ := template.ParseFiles("index.html")
		errorValue := map[string]interface{}{
			"error": true,
		}
		t.Execute(w, errorValue)
	} else { // When city was found
		// Read content of html file and returns a Template
		t, _ := template.ParseFiles("weather.html")
		// Execute the template, writing the generated HTML to the http.ResponseWriter
		t.Execute(w, cityValues)
	}

	elapsed := time.Since(start)
	log.Printf("Index page took %s", elapsed)
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
