package repository

import (
	"delivery/internal/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *model.Order) (*model.Order, error)
}

type GormOrderRepository struct {
	DB *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{DB: db}
}

func (r *GormOrderRepository) CreateOrder(order *model.Order) (*model.Order, error) {
	err := r.DB.Create(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}
