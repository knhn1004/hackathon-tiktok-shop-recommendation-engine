// File: models/recommendation.go

package models

import (
	"time"

	"github.com/lib/pq"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)
   type ArticleEmbedding struct {
       ArticleID uint `gorm:"primaryKey"`
       Article   Article
       Embedding pgvector.Vector `gorm:"type:vector(384)"`
   }

   type ProductEmbedding struct {
       ProductID uint `gorm:"primaryKey"`
       Product   Product
       Embedding pgvector.Vector `gorm:"type:vector(384)"`
   }

type UserArticleInteraction struct {
	gorm.Model
	UserProfileID    uint
	UserProfile      UserProfile
	ArticleID        uint
	Article          Article
	InteractionType  string
	Duration         int
}

type UserProductInteraction struct {
	gorm.Model
	UserProfileID    uint
	UserProfile      UserProfile
	ProductID        uint
	Product          Product
	InteractionType  string
}

type UserArticleRecommendation struct {
       UserProfileID       uint `gorm:"primaryKey"`
       UserProfile         UserProfile
       RecommendedArticles pq.Int64Array `gorm:"type:integer[]"`
       LastUpdated         time.Time
   }

   type UserProductRecommendation struct {
       UserProfileID       uint `gorm:"primaryKey"`
       UserProfile         UserProfile
       RecommendedProducts pq.Int64Array `gorm:"type:integer[]"`
       LastUpdated         time.Time
   }
