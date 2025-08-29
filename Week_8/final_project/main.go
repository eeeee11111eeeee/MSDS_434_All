package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// RealEstate defines a structure for a real estate listing.
type RealEstate struct {
	ID                       int     `json:"id"`
	PrefectureName           string  `json:"prefecture_name,omitempty"`
	CityName                 string  `json:"city_name,omitempty"`
	DistrictName             string  `json:"district_name,omitempty"`
	NearestStationName       string  `json:"nearest_station_name,omitempty"`
	DistanceToStationMinutes float64 `json:"distance_to_station_minutes,omitempty"`
	Layout                   string  `json:"layout,omitempty"`
	AreaSqm                  int     `json:"area_sqm,omitempty"`
	ConstructionYear         float64 `json:"construction_year,omitempty"`
	BuildingStructure        string  `json:"building_structure,omitempty"`
	CityPlanning             string  `json:"city_planning,omitempty"`
	BuildingCoverageRatio    float64 `json:"building_coverage_ratio,omitempty"`
	FloorAreaRatio           float64 `json:"floor_area_ratio,omitempty"`
	TransactionDate          string  `json:"transaction_date,omitempty"`
}

// GetListingsHandler is used to get all listings from the JSON file.
func GetListingsHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read JSON file using os.ReadFile
		data, err := os.ReadFile("./data/data.json")
		if err != nil {
			log.Printf("File read error: %v", err)
			http.Error(rw, "Failed to read data file", http.StatusInternalServerError)
			return
		}

		// Write the body with JSON data
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(data)
	}
}

// AddListingHandler is used to add a new listing to the JSON file.
func AddListingHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read incoming JSON from request body using io.ReadAll
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		var newListing RealEstate
		err = json.Unmarshal(data, &newListing)
		if err != nil {
			http.Error(rw, "Invalid Data Format", http.StatusExpectationFailed)
			return
		}

		// Load existing listings from the file
		fileData, err := os.ReadFile("./data/data.json")
		if err != nil {
			log.Printf("File read error: %v", err)
			http.Error(rw, "Failed to read data file", http.StatusInternalServerError)
			return
		}

		// The file is an array of objects, so we unmarshal into a slice
		var listings []RealEstate
		err = json.Unmarshal(fileData, &listings)
		if err != nil {
			// If the file is empty, initialize an empty slice
			if string(fileData) == "" {
				listings = []RealEstate{}
			} else {
				log.Printf("JSON unmarshal error: %v", err)
				http.Error(rw, "Failed to parse data file", http.StatusInternalServerError)
				return
			}
		}

		// Add the new listing to our list
		listings = append(listings, newListing)

		// Write the updated list back to the JSON file using os.WriteFile
		updatedData, err := json.MarshalIndent(listings, "", "    ")
		if err != nil {
			log.Printf("JSON marshal error: %v", err)
			http.Error(rw, "Failed to marshal data", http.StatusInternalServerError)
			return
		}
		err = os.WriteFile("./data/data.json", updatedData, os.ModePerm)
		if err != nil {
			log.Printf("File write error: %v", err)
			http.Error(rw, "Failed to write data to file", http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Real Estate Listing"))
	}
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Route properly to respective handlers
	router.Handle("/listings", GetListingsHandler()).Methods("GET")
	router.Handle("/listings", AddListingHandler()).Methods("POST")

	// Create a new server and assign the router
	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}

	fmt.Println("Starting Real Estate server on Port 9090")
	// Start the server on the defined port
	log.Fatal(server.ListenAndServe())
}
