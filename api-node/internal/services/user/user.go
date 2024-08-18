// File: internal/services/user/user.go

package user

import (
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
)

// CreateUser creates a new user profile
func CreateUser(userProfile *models.UserProfile) error {
	return db.DB.Create(userProfile).Error
}

// GetUserByID retrieves a user profile by ID
func GetUserByID(userID string) (*models.UserProfile, error) {
	var user models.UserProfile
	err := db.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user profile
func UpdateUser(userProfile *models.UserProfile) error {
	return db.DB.Save(userProfile).Error
}

// DeleteUser deletes a user profile by ID
func DeleteUser(userID string) error {
	return db.DB.Where("user_id = ?", userID).Delete(&models.UserProfile{}).Error
}
