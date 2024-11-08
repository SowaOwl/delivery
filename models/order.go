package models

import "time"

type Order struct {
	ID           uint `gorm:"primaryKey"`
	OrderTime    time.Time
	OrderEndTime time.Time
	UserID       uint `gorm:"index;foreignKey;not null"`
	ChefID       uint `gorm:"index;foreignKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	User   User   `gorm:"foreignKey:UserID"`
	Chef   Chef   `gorm:"foreignKey:ChefID"`
	Dishes []Dish `gorm:"many2many:order_dishes"`
}
