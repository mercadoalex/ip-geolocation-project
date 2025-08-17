package main

import (
	"encoding/json"
	"ip-geolocation-project/src/geo"
	"ip-geolocation-project/src/utils"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the logger for info, warning, and error messages
	logger := utils.NewLogger()

	// Open the GeoLite2 database (update path if needed)
	db, err := geo.NewGeoLite2DB("maxmind/GeoLite2-City.mmdb")
	if err != nil {
		logger.Error("Failed to open GeoLite2 database: " + err.Error())
		return
	}
	defer db.Close()

	// Create a new router using gorilla/mux
	router := mux.NewRouter()

	// Define the IP geolocation endpoint
	// Example: GET /geolocation/8.8.8.8
	router.HandleFunc("/geolocation/{ip}", func(w http.ResponseWriter, r *http.Request) {
		ip := mux.Vars(r)["ip"] // Extract IP from URL
		location, err := db.GetLocation(ip)
		if err != nil {
			logger.Error("Failed to get location: " + err.Error())
			http.Error(w, "Unable to retrieve location", http.StatusInternalServerError)
			return
		}

		// Respond with the location data as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(location)
	}).Methods("GET")

	// Set up CORS middleware to allow requests from any origin
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),            // Allow all origins
		handlers.AllowedMethods([]string{"GET"}),          // Allow GET requests
		handlers.AllowedHeaders([]string{"Content-Type"}), // Allow Content-Type header
	)(router)

	// Start the server with CORS enabled and log its status
	logger.Info("Starting server on :8080 with CORS enabled")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		logger.Error("Server failed to start: " + err.Error())
	}
}
