package repository

import (
	"delivery/internal/model"
	"gorm.io/gorm"
)

type ChefRepository interface {
	GetAllSortedByOrders() ([]model.Chef, error)
}

type GormChefRepository struct {
	DB *gorm.DB
}

func NewGormChefRepository(DB *gorm.DB) *GormChefRepository {
	return &GormChefRepository{DB: DB}
}

func (r *GormChefRepository) GetAllSortedByOrders() ([]model.Chef, error) {
	var chefs []model.Chef

	result := r.DB.
		Preload("Orders").
		Joins("LEFT JOIN orders ON orders.chef_id = chefs.id").
		Group("chefs.id").
		Order("COUNT(orders.id) DESC").
		Find(&chefs)

	if result.Error != nil {
		return nil, result.Error
	}

	return chefs, nil
}
