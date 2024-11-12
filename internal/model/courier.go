package model

import "time"

type Courier struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index;foreignKey;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}
