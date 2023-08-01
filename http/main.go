package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("Content-Type", "Application-json")
		response.Write([]byte("Hello World!"))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
