package main

import (
	"log"
	"github.com/festusfrankzin/ecommerceApi/cmd/api"
	"github.com/festusfrankzin/ecommerceApi/db"
)

func main() {


	// Running the database
	db, err :=  db.ConnectDB()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer db.Close()


	// Running the API server
	server := api.CreateApiServer(":8000", nil)
	if err := server.Run(); err != nil {
		log.Fatal("Error creating api server")
	}
}

