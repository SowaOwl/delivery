package model

import "time"

type User struct {
	ID         uint   `gorm:"primaryKey"`
	FirstName  string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	MiddleName string
	Email      string `gorm:"not null"`
	Password   string `gorm:"not null"`
	BirthDate  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Orders []Order `gorm:"foreignKey:UserID"`
}
