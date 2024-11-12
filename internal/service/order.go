package service

import (
	"delivery/internal/constants"
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

func OrderToWaitDelivery(order model.Order, db *gorm.DB) error {
	orderRepo := repository.NewGormOrderRepository(db)

	err := orderRepo.UpdateStatus(order.ID, constants.WaitDeliver)
	if err != nil {
		return err
	}

	return nil
}

func OrderToOnTheWay(order model.Order, db *gorm.DB) error {
	orderRepo := repository.NewGormOrderRepository(db)

	err := orderRepo.UpdateStatus(order.ID, constants.OnTheWay)
	if err != nil {
		return err
	}

	return nil
}

func OrderToDelivered(order model.Order, db *gorm.DB) error {
	orderRepo := repository.NewGormOrderRepository(db)

	err := orderRepo.UpdateStatus(order.ID, constants.Delivered)
	if err != nil {
		return err
	}

	return nil
}

func CancelledOrder(order model.Order, db *gorm.DB) error {
	orderRepo := repository.NewGormOrderRepository(db)

	err := orderRepo.UpdateStatus(order.ID, constants.Cancelled)
	if err != nil {
		return err
	}

	return nil
}
