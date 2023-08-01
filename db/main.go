package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	urlConnection := "user=lucas dbname=devbook sslmode=disable"
	db, err := sql.Open("postgres", urlConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	println(rows)
}
