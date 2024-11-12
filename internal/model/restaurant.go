package model

import "time"

type Restaurant struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Chefs []Chef `gorm:"foreignKey:RestaurantID"`
}
