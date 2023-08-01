package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name  string `json:"name"`
	Age   uint   `json:"agr"`
	Breed string `json:"breed"`
}

func main() {
	dog := Dog{"Black", 6, "SRN"}
	convertedJson, err := json.Marshal(dog)
	if err != nil {
		log.Fatal(err)
	}
	output := bytes.NewBuffer(convertedJson)
	fmt.Println(output)
	fmt.Println("--------------------------------")
	dog2 := map[string]string{
		"name":  "Zé",
		"age":   "2",
		"breed": "Rotweiller",
	}
	dog2Json, err := json.Marshal(dog2)
	if err != nil {
		log.Fatal(err)
	}
	output2 := bytes.NewBuffer(dog2Json)
	fmt.Println(output2)

}
