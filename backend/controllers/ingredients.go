package controllers

import (
	"calorie/data"
	"net/http"

	"github.com/gorilla/mux"
)

func ingredients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	Data := data.Data

	
}