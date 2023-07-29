package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Breed string `json:"breed"`
}

func main() {
	dogJson := `{"age":2,"breed":"Rotweiller","name":"Zé"}`
	var dog Dog
	if err := json.Unmarshal([]byte(dogJson), &dog); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dog)
}
