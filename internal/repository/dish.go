package repository

import (
	"delivery/internal/model"
	"errors"
	"gorm.io/gorm"
)

type DishRepository interface {
	CreateOrUpdateDish(dish *model.Dish) error
	GetAll() ([]*model.Dish, error)
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

	return r.DB.Model(&existingDish).Updates(dish).Error
}

func (r *GormDishRepository) GetAll() ([]*model.Dish, error) {
	var dishes []*model.Dish

	err := r.DB.Find(&dishes).Error
	if err != nil {
		return nil, err
	}

	return dishes, nil
}
