package repository

import (
	"delivery/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
}

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		DB: db,
	}
}

func (r *GormUserRepository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}

	err := r.DB.First(user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
