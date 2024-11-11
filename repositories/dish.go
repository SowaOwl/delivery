package repositories

import (
	"delivery/models"
	"errors"
	"fmt"
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

	err := r.DB.First(&existingDish, dish.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = r.DB.Create(dish).Error
		fmt.Printf(err.Error())
	}

	if err != nil {
		return err
	}

	existingDish.ID = dish.ID
	existingDish.Name = dish.Name
	existingDish.Description = dish.Description
	existingDish.Price = dish.Price
	return r.DB.Save(&existingDish).Error
}
