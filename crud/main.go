package main

import (
	usersControllers "crud/controllers/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", usersControllers.Create).Methods(http.MethodPost)
	println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
