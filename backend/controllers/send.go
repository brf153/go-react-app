package controllers

import (
	"encoding/json"
	"net/http"
	"calorie"
)

func SendData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	data := main.Data
	jsonData, err := json.Marshal(data)
	if err != nil{
		http.Error(w,"Data mei error",http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}