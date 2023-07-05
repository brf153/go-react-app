package controllers

import (
	"calorie/data"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Set the CORS headers to allow requests from any origin.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	Data := &data.Data

	for index, dish := range *Data {
		if dish.ID == id {
			*Data = append((*Data)[:index], (*Data)[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(Data)
}
