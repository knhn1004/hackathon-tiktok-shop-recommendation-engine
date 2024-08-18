// File: models/shop.go

package models

import (
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	CreatorID uint `gorm:"constraint:OnDelete:CASCADE;"`
	Creator     Creator
	Name        string
	Description string
}

type Category struct {
	gorm.Model
	Name     string
	ParentID *uint
	Parent   *Category
}

type Product struct {
	gorm.Model
	Shop        Shop
	ShopID      uint `gorm:"constraint:OnDelete:CASCADE;"`
	CategoryID  uint `gorm:"foreignKey:ID;constraint:OnDelete:SET NULL;"`
	Category    Category
	Title       string
	Description string
	Price       float64
	ImageURL    string
}