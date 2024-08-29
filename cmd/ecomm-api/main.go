package main

import (
	"log"

	"github.com/abhishek622/go-sqlx/db"
	"github.com/abhishek622/go-sqlx/ecomm-api/handler"
	"github.com/abhishek622/go-sqlx/ecomm-api/server"
	"github.com/abhishek622/go-sqlx/ecomm-api/storer"
	"github.com/ianschenck/envflag"
)

const minSecretKeySize = 32

func main() {
	var secretKey = envflag.String("SECRET_KEY", "01234567890123456789012345678901", "secret key for JWT signing")
	if len(*secretKey) < minSecretKeySize {
		log.Fatal("SECRET_KEY must be at least %d characters", minSecretKeySize)
	}

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error opening database: %v", err)
	}

	defer db.Close()
	log.Println("Successfully connected to DB....")

	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv, *secretKey)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
