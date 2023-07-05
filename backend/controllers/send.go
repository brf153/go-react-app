package controllers

import (
	"calorie/data"
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, r *http.Request) {
	// Set the CORS headers to allow requests from any origin.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")

	data := data.Data
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Data mei error", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
