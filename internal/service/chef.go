package service

import (
	"delivery/internal/model"
	"delivery/internal/repository"
	"errors"
	"gorm.io/gorm"
)

func GetFreeChef(db *gorm.DB) (*model.Chef, error) {
	repo := repository.NewGormChefRepository(db)

	chefs, err := repo.GetAllSortedByOrders()
	if err != nil {
		return nil, err
	}

	for _, chef := range chefs {
		if len(chef.Orders) < int(chef.MaxOrders) {
			return &chef, nil
		}
	}

	return nil, errors.New("not found free chef pleas try later")
}
