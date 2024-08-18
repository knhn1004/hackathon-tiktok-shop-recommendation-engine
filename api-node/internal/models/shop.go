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

type Order struct {
	gorm.Model
	UserID    string
	ShopID    uint
	Status    string // e.g., "pending", "completed", "cancelled"
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	Order     Order   `gorm:"foreignKey:OrderID"`
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int
	Price     float64 // Price at the time of order
}

type OrderItemInput struct {
	ProductID uint `json:"productId"`
	Quantity  int  `json:"quantity"`
}