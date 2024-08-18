// File: models/user.go

package models

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserID     string `json:"userId" gorm:"uniqueIndex"`
	Username   string `json:"username"`
	Bio        string `json:"bio"`
	AvatarURL  string `json:"avatarUrl"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Creator struct {
    gorm.Model   `json:"-"`
    ID           uint        `json:"id" gorm:"primaryKey"`
    UserProfileID uint        `json:"userProfileId"`
    UserProfile   UserProfile `json:"userProfile" gorm:"foreignKey:UserProfileID;constraint:OnDelete:CASCADE;"`
    CreatedAt     time.Time   `json:"createdAt"`
    UpdatedAt     time.Time   `json:"updatedAt"`
}

func GetUserProfileByUserID(db *gorm.DB, userID string) (*UserProfile, error) {
	var userProfile UserProfile
	result := db.Where("user_id = ?", userID).First(&userProfile)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userProfile, nil
}