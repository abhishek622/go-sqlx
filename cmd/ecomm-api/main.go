package main

import (
	"log"

	"github.com/abhishek622/go-sqlx/db"
	"github.com/abhishek622/go-sqlx/ecomm-api/server"
	"github.com/abhishek622/go-sqlx/ecomm-api/storer"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error opening database: %v", err)
	}

	defer db.Close()
	log.Println("Successfully connected to DB....")

	st := storer.NewMySQLStorer(db.GetDB())
	_ = server.NewServer(st)
}
