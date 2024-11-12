package model

import "time"

type Chef struct {
	ID           uint `gorm:"primaryKey"`
	MaxOrders    uint `gorm:"default:1"`
	UserID       uint `gorm:"index;foreignKey;not null"`
	RestaurantID uint `gorm:"index;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	User       User       `gorm:"foreignKey:UserID"`
	Restaurant Restaurant `gorm:"foreignKey:RestaurantID"`
	Orders     []Order    `gorm:"foreignKey:ChefID"`
}
