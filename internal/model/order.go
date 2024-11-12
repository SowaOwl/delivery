package model

import "time"

type Order struct {
	ID           uint `gorm:"primaryKey"`
	OrderTime    time.Time
	OrderEndTime *time.Time
	Status       uint    `gorm:"default:1"`
	OrderSum     float32 `gorm:"type:decimal(10,2)"`
	UserID       uint    `gorm:"index;foreignKey;not null"`
	ChefID       uint    `gorm:"index;foreignKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	User   User   `gorm:"foreignKey:UserID"`
	Chef   Chef   `gorm:"foreignKey:ChefID"`
	Dishes []Dish `gorm:"many2many:order_dishes"`
}
