package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Conn() (*sql.DB, error) {
	connectionString := "user=lucas dbname=devbook sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
