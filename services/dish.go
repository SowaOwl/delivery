package services

import (
	"delivery/providers"
	"delivery/repositories"
	"gorm.io/gorm"
)

func UpdateDishes(provider providers.DishProvider, db *gorm.DB) error {
	repo := repositories.NewGormDishRepository(db)

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
