package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	datasources "weatherHTTP/datasources/NationalWeatherService"
	"weatherHTTP/objects"
)

// basic http server that exposes one endpoint, /getWeatherData, and takes an object of type objects.WeatherInput as a request body
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getWeatherData", getweatherDataHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server on port 8080
	log.Println("Listening at: localhost:8080!")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func getweatherDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req objects.WeatherInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	weatherData, err := datasources.GetWeatherData(r.Context(), req.Latitude, req.Longitude)
	if err != nil {
		log.Printf("error getting weather data: %v", err)
		http.Error(w, "Failed to retrieve weather data", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(weatherData); err != nil {
		log.Printf("error encoding weather response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
