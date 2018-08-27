package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	// Read content of html file and returns a Template
	t, _ := template.ParseFiles("index.html")
	// Execute the template, writing the generated HTML to the http.ResponseWriter
	t.Execute(w, nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieves value from form
	cityName := r.FormValue("city")
	apiKey := os.Getenv("API_ROUNDS")

	if apiKey == "" {
		apiKey = "f650d6a017adae74d1aba918770e6389"
	}

	// Build url
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?APPID=%s&q=%s&mode=json&units=metric", apiKey, cityName)

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
