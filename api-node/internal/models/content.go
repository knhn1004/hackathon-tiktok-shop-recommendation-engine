// File: models/content.go

package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	CreatorID uint `gorm:"constraint:OnDelete:CASCADE;"`
	Creator    Creator
	Title      string
	Content string
	Likes []ArticleLike `gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE;"`
	Views      int
	Tags       []Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	gorm.Model
	Name string
}

type Comment struct {
	gorm.Model
	UserProfile   UserProfile
	Article       Article
	Content       string
	UserProfileID uint `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;"`
	ArticleID     uint `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;"`
}

type ArticleLike struct {
	UserProfileID uint      `gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	ArticleID     uint      `gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	CreatedAt     time.Time
}