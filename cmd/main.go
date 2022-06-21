package main

import (
	"e-comerce/app"
	"log"
	"net/http"
)

func main() {
	router := app.Initialize()
	log.Println("Server run is 9000 port...")
	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatal(err)
	}
}
