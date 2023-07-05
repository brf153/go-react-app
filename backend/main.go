package main

import (
	"calorie/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			// Handle preflight requests. Set the status code and return.
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}

func main() {
	PORT := ":4000"

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.SendData).Methods("GET")

	r.HandleFunc("/dish/{id}", controllers.DeleteHandler).Methods("DELETE")

	r.HandleFunc("/dish/ingredients/{id}", controllers.Ingredients)

	http.Handle("/", CORSMiddleware(r))

	fmt.Printf("Server started on port %s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
