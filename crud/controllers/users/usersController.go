package usersControllers

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Create(response http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	var user user
	if err = json.Unmarshal(body, &user); err != nil {
		response.WriteHeader(401)
		return
	}
	fmt.Println(user)
	database, err := db.Conn()
	if err != nil {
		log.Fatal(err)
		response.WriteHeader(500)
		return
	}
	defer database.Close()
	statement, err := database.Prepare("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id")
	if err != nil {
		response.WriteHeader(500)
		return
	}
	defer statement.Close()
	var id string
	err = statement.QueryRow(user.Name, user.Email).Scan(&id)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	user.ID = id
	output, err := json.Marshal(user)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	response.WriteHeader(http.StatusCreated)
	response.Write(output)
}
