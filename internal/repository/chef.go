package repository

import (
	"delivery/internal/constants"
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
		Joins("LEFT JOIN orders ON orders.chef_id = chefs.id AND orders.status IN (?, ?)", uint(constants.New), uint(constants.OnKitchen)).
		Select("chefs.*, COUNT(orders.id) AS order_count").
		Group("chefs.id").
		Order("order_count").
		Find(&chefs)

	if result.Error != nil {
		return nil, result.Error
	}

	return chefs, nil
}
