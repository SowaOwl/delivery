package service

import (
	"delivery/internal/model"
	"delivery/internal/repository"
	"gorm.io/gorm"
	"time"
)

type CreateOrderDTO struct {
	Status   uint
	UserId   uint
	Dishes   []model.Dish
	OrderSum float32
}

func CreateOrder(orderDTO CreateOrderDTO, db *gorm.DB) error {
	orderRepo := repository.NewGormOrderRepository(db)

	chef, err := GetFreeChef(db)
	if err != nil {
		return err
	}

	newOrder := model.Order{
		OrderTime: time.Now(),
		UserID:    orderDTO.UserId,
		Chef:      *chef,
		Dishes:    orderDTO.Dishes,
		OrderSum:  orderDTO.OrderSum,
	}

	_, err = orderRepo.CreateOrder(&newOrder)
	if err != nil {
		return err
	}

	return nil
}
