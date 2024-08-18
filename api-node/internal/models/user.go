// File: models/user.go

package models

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	UserID    string `gorm:"uniqueIndex"`
	Username  string
	Bio       string
	AvatarURL string
}

type Creator struct {
	gorm.Model
	UserProfileID uint
	UserProfile   UserProfile
}
