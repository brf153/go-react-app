package controllers

import (
	"calorie/data"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Ingredients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var updatedDish data.DataStruct
	if err := json.Unmarshal(body, &updatedDish); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	w.Write(body)

	vars := mux.Vars(r)
	id := vars["id"]

	Data := &data.Data

	for _,dish:= range *Data{
		if dish.ID==id{
			dish.Ingredients = updatedDish.Ingredients
		}
	}

	json.NewEncoder(w).Encode(Data)
}
