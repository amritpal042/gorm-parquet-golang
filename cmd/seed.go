package main

import (
	"log"

	"gorm-parquet-golang/internal/db"
	"gorm-parquet-golang/internal/mockdata"
)

func main() {
	// Set up the database connection
	connection, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Seed mock data into the database
	mockdata.SeedData(connection)

	log.Println("Seeding complete.")
}
