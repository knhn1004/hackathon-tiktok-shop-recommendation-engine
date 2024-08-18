// File: models/shop.go

package models

import (
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	CreatorID   uint
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
	ShopID      uint
	Shop        Shop
	CategoryID  uint
	Category    Category
	Title       string
	Description string
	Price       float64
	ImageURL    string
}