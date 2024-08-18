// File: models/recommendation.go

package models

import (
	"time"

	"github.com/lib/pq"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type ArticleEmbedding struct {
    ArticleID uint            `json:"articleId" gorm:"primaryKey"`
    Article   Article         `json:"article"`
    Embedding pgvector.Vector `json:"embedding" gorm:"type:vector(384)"`
}

type ProductEmbedding struct {
    ProductID uint            `json:"productId" gorm:"primaryKey"`
    Product   Product         `json:"product"`
    Embedding pgvector.Vector `json:"embedding" gorm:"type:vector(384)"`
}

type UserArticleInteraction struct {
    gorm.Model    `json:"-"`
    ID            uint        `json:"id" gorm:"primaryKey"`
    UserProfileID uint        `json:"userProfileId"`
    UserProfile   UserProfile `json:"userProfile"`
    ArticleID     uint        `json:"articleId"`
    Article       Article     `json:"article"`
    InteractionType string    `json:"interactionType"`
    Duration      int         `json:"duration"`
    CreatedAt     time.Time   `json:"createdAt"`
    UpdatedAt     time.Time   `json:"updatedAt"`
}

type UserProductInteraction struct {
    gorm.Model    `json:"-"`
    ID            uint        `json:"id" gorm:"primaryKey"`
    UserProfileID uint        `json:"userProfileId"`
    UserProfile   UserProfile `json:"userProfile"`
    ProductID     uint        `json:"productId"`
    Product       Product     `json:"product"`
    InteractionType string    `json:"interactionType"`
    CreatedAt     time.Time   `json:"createdAt"`
    UpdatedAt     time.Time   `json:"updatedAt"`
}

type UserArticleRecommendation struct {
    UserProfileID       uint           `json:"userProfileId" gorm:"primaryKey"`
    UserProfile         UserProfile    `json:"userProfile"`
    RecommendedArticles pq.Int64Array  `json:"recommendedArticles" gorm:"type:integer[]"`
    LastUpdated         time.Time      `json:"lastUpdated"`
}

type UserProductRecommendation struct {
    UserProfileID       uint           `json:"userProfileId" gorm:"primaryKey"`
    UserProfile         UserProfile    `json:"userProfile"`
    RecommendedProducts pq.Int64Array  `json:"recommendedProducts" gorm:"type:integer[]"`
    LastUpdated         time.Time      `json:"lastUpdated"`
}