package main

import (
	"fmt"
	"log"
	"net/http"
	"calorie/controllers"
)

func main() {
	PORT := ":4000"

	http.HandleFunc("/",controllers.SendData)

	http.HandleFunc("/dish/:id",controllers.DeleteHandler)

	// http.Get("/",helloHandler)

	fmt.Printf("Server started on port %s", PORT)
	err := http.ListenAndServe(PORT,nil)
	if err!=nil{
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "The server is working!")
}