package main

import (
	"log"

	"github.com/abhishek622/go-sqlx/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error opening database: %v", err)
	}

	defer db.Close()
	log.Println("Successfully connected to DB....")
}
