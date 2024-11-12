package repository

import (
	"delivery/internal/model"
	"errors"
	"gorm.io/gorm"
)

type DishRepository interface {
	CreateOrUpdateDish(dish *model.Dish) error
}

type GormDishRepository struct {
	DB *gorm.DB
}

func NewGormDishRepository(db *gorm.DB) *GormDishRepository {
	return &GormDishRepository{DB: db}
}

func (r *GormDishRepository) CreateOrUpdateDish(dish *model.Dish) error {
	var existingDish model.Dish

	err := r.DB.First(&existingDish, "id = ?", dish.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = r.DB.Create(&dish).Error
		if err != nil {
			return err
		}

		return nil
	}

	if err != nil {
		return err
	}

	return r.DB.Updates(&dish).Error
}