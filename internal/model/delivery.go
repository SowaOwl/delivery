package model

import "time"

type Delivery struct {
	ID           uint `gorm:"primaryKey"`
	CourierID    uint `gorm:"index"`
	DeliveryTime time.Time
	Address      string `gorm:"not null"`
	Floor        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
