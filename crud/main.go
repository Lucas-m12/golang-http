package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
