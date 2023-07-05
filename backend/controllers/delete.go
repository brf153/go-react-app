package controllers

import (
	"calorie/data"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	Data := data.Data

	for index, dish := range Data {
		if dish.ID == id {
			Data = append(Data[:index], Data[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(Data)
}
