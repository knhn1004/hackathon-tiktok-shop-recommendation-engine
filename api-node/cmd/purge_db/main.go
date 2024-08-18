package main

import (
	"log"

	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/config"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	err = db.InitDB(
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Add logic here to purge the database
	// For example, you might drop all tables or truncate them

	log.Println("Database purged successfully")
}