// File: models/content.go

package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	CreatorID  uint
	Creator    Creator
	Title      string
	ContentURL string
	Likes      int
	Views      int
	Tags       []Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	gorm.Model
	Name string
}

type Comment struct {
	gorm.Model
	UserProfileID uint
	UserProfile   UserProfile
	ArticleID     uint
	Article       Article
	Content       string
}

type ArticleLike struct {
	UserProfileID uint
	ArticleID     uint
	CreatedAt     time.Time
}