package main

import (
	"log"

	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/config"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
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

	// Purge the database by dropping all tables
	err = purgeDatabase()
	if err != nil {
		log.Fatalf("Failed to purge database: %v", err)
	}

	log.Println("Database purged successfully")
}

func purgeDatabase() error {
	// List of all models to be dropped
	models := []interface{}{
		&models.UserProfile{},
		&models.Creator{},
		&models.Article{},
		&models.Tag{},
		&models.Shop{},
		&models.Category{},
		&models.Product{},
		&models.Comment{},
		&models.ArticleLike{},
		&models.ArticleEmbedding{},
		&models.ProductEmbedding{},
		&models.UserArticleInteraction{},
		&models.UserProductInteraction{},
		&models.UserArticleRecommendation{},
		&models.UserProductRecommendation{},
		&models.KafkaEvent{},
	}

	// Drop all tables
	for _, model := range models {
		if err := db.DB.Migrator().DropTable(model); err != nil {
			return err
		}
	}
	// Auto migrate all models
	for _, model := range models {
		if err := db.DB.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}