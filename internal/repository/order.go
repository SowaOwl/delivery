package repository

import (
	"delivery/internal/constants"
	"delivery/internal/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *model.Order) (*model.Order, error)
	UpdateStatus(orderId uint, status constants.OrderStatus) error
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

func (r *GormOrderRepository) UpdateStatus(orderId uint, status constants.OrderStatus) error {
	order := &model.Order{}

	err := r.DB.First(&order, orderId).Error
	if err != nil {
		return err
	}

	order.Status = status.ToUInt()

	return r.DB.Save(&order).Error
}
