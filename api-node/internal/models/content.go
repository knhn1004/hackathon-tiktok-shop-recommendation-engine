// File: models/content.go

package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatorID  uint      `json:"creatorId" gorm:"constraint:OnDelete:CASCADE;"`
	Creator    Creator   `json:"creator" gorm:"foreignKey:CreatorID"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Likes      []ArticleLike `json:"likes" gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE;"`
	Views      int       `json:"views"`
	Tags       []Tag     `json:"tags" gorm:"many2many:article_tags;"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Tag struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Comment struct {
	gorm.Model    `json:"-"`
	ID            uint        `json:"id" gorm:"primaryKey"`
	UserProfile   UserProfile `json:"userProfile"`
	Article       Article     `json:"article"`
	Content       string      `json:"content"`
	UserProfileID uint        `json:"userProfileId" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;"`
	ArticleID     uint        `json:"articleId" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}

type ArticleLike struct {
	UserProfileID uint      `json:"userProfileId" gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	ArticleID     uint      `json:"articleId" gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	CreatedAt     time.Time `json:"createdAt"`
}