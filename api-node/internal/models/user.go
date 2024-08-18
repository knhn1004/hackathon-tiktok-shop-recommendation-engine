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
	UserProfile   UserProfile `gorm:"foreignKey:UserProfileID;constraint:OnDelete:CASCADE;"`
}

func GetUserProfileByUserID(db *gorm.DB, userID string) (*UserProfile, error) {
	var userProfile UserProfile
	result := db.Where("user_id = ?", userID).First(&userProfile)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userProfile, nil
}
