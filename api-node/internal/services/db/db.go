// File: internal/services/db/db.go

package db

import (
	"fmt"
	"log"

	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(host, user, password, dbname string, port int) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	err = migrateSchema()
	if err != nil {
		return fmt.Errorf("failed to migrate schema: %v", err)
	}

	log.Println("Database connected and migrated successfully")
	return nil
}

// migrateSchema automatically migrates the schema
func migrateSchema() error {
	return DB.AutoMigrate(
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
	)
}

// CreateUser creates a new user profile
func CreateUser(user *models.UserProfile) error {
	return DB.Create(user).Error
}

// GetUser retrieves a user profile by ID
func GetUser(id uint) (*models.UserProfile, error) {
	var user models.UserProfile
	err := DB.First(&user, id).Error
	return &user, err
}

// CreateArticle creates a new article
func CreateArticle(article *models.Article) error {
	return DB.Create(article).Error
}

// GetArticle retrieves an article by ID
func GetArticle(id uint) (*models.Article, error) {
	var article models.Article
	err := DB.First(&article, id).Error
	return &article, err
}

// Add more database operations as needed...