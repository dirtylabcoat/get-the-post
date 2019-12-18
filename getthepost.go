package main

import (
	"encoding/json"
	"net/http"
)

type requestData struct {
	URL    string `json:"url"`
	Method string `json:"method"`
	Data   string `json:"data"`
}

// GetThePost is a function that proxies POST, PUT and DELETE calls over GET
func GetThePost(w http.ResponseWriter, r *http.Request) {
	input := requestData{}

	// Check for non-empty/nil input
	// Check for existence of actual values in url, method and data
	// Validate size of data, max 3K
	// Only allowed values for method are POST, PUT or DELETE
	// Is url an URL?

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(input)
	}
}

