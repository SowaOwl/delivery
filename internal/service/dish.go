package service

import (
	"delivery/internal/provider"
	"delivery/internal/repository"
	"gorm.io/gorm"
)

func UpdateDishes(provider provider.DishProvider, db *gorm.DB) error {
	repo := repository.NewGormDishRepository(db)

	dishes, err := provider.GetDishes()
	if err != nil {
		return err
	}

	for _, dish := range dishes {
		err := repo.CreateOrUpdateDish(&dish)
		if err != nil {
			return err
		}
	}

	return nil
}
