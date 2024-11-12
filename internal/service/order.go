package service

import (
	"delivery/internal/model"
	"delivery/internal/repository"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type CreateOrderDTO struct {
	Status   uint
	UserId   uint
	Dishes   []uint
	OrderSum float32
}

func CreateOrder(orderDTO CreateOrderDTO, db *gorm.DB) {
	repo := repository.NewGormOrderRepository(db)

	newOrder := model.Order{
		OrderTime: time.Now(),
		UserID:    orderDTO.UserId,
		ChefID:    1,
	}

	_, err := repo.CreateOrder(&newOrder)
	if err != nil {
		fmt.Println(err)
	}
}
