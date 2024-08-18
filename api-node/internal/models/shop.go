// File: models/shop.go

package models

import (
	"time"

	"gorm.io/gorm"
)

type Shop struct {
    gorm.Model `json:"-"`
    ID          uint   `json:"id" gorm:"primaryKey"`
    CreatorID   uint   `json:"creatorId" gorm:"constraint:OnDelete:CASCADE;"`
    Creator     Creator `json:"creator"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}

type Category struct {
    gorm.Model `json:"-"`
    ID       uint      `json:"id" gorm:"primaryKey"`
    Name     string    `json:"name"`
    ParentID *uint     `json:"parentId"`
    Parent   *Category `json:"parent"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type Product struct {
    gorm.Model `json:"-"`
    ID          uint     `json:"id" gorm:"primaryKey"`
    Shop        Shop     `json:"shop" gorm:"foreignKey:ShopID"`
    ShopID      uint     `json:"shopId" gorm:"constraint:OnDelete:CASCADE;"`
    Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
    CategoryID  uint     `json:"categoryId" gorm:"constraint:OnDelete:SET NULL;"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Price       float64  `json:"price"`
    ImageURL    string   `json:"imageUrl"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}

type Order struct {
    gorm.Model `json:"-"`
    ID         uint        `json:"id" gorm:"primaryKey"`
    UserID     string      `json:"userId"`
    ShopID     uint        `json:"shopId"`
    Status     string      `json:"status"`
    OrderItems []OrderItem `json:"orderItems"`
    CreatedAt  time.Time   `json:"createdAt"`
    UpdatedAt  time.Time   `json:"updatedAt"`
}

type OrderItem struct {
    gorm.Model `json:"-"`
    ID        uint    `json:"id" gorm:"primaryKey"`
    OrderID   uint    `json:"orderId"`
    Order     Order   `json:"order" gorm:"foreignKey:OrderID"`
    ProductID uint    `json:"productId"`
    Product   Product `json:"product" gorm:"foreignKey:ProductID"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type OrderItemInput struct {
	ProductID uint `json:"productId"`
	Quantity  int  `json:"quantity"`
}