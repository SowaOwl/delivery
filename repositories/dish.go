package repositories

import (
	"delivery/models"
	"errors"
	"gorm.io/gorm"
)

type DishRepository interface {
	CreateOrUpdateDish(dish *models.Dish) error
}

type GormDishRepository struct {
	DB *gorm.DB
}

func NewGormDishRepository(db *gorm.DB) *GormDishRepository {
	return &GormDishRepository{DB: db}
}

func (r *GormDishRepository) CreateOrUpdateDish(dish *models.Dish) error {
	var existingDish models.Dish

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
