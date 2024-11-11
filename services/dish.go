package services

import (
	"delivery/models"
	"delivery/providers"
	"delivery/repositories"
	"gorm.io/gorm"
)

func UpdateDishes(provider providers.DishProvider, db *gorm.DB) error {
	repo := repositories.NewGormDishRepository(db)

	//dishes, err := provider.GetDishes()
	//if err != nil {
	//	return err
	//}
	//
	//for _, dish := range dishes {
	//	err := repo.CreateOrUpdateDish(&dish)
	//	if err != nil {
	//		return err
	//	}
	//}

	dish := models.Dish{
		ID:          1,
		Name:        "TEST",
		Description: "TEST",
		Price:       100.12,
	}

	err := repo.CreateOrUpdateDish(&dish)

	if err != nil {
		return err
	}

	return nil
}
