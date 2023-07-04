package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":4000"

	err := http.ListenAndServe(PORT,nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("Server started on port %s", PORT)
}
