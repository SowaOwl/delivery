package service

import (
	"delivery/internal/model"
	"delivery/internal/provider"
	"delivery/internal/repository"
	"gorm.io/gorm"
	"sync"
)

func UpdateDishes(provider provider.DishProvider, db *gorm.DB) error {
	repo := repository.NewGormDishRepository(db)

	dishes, err := provider.GetDishes()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	workersNum := 20

	ch := make(chan struct{}, workersNum)

	for i := 0; i < workersNum; i++ {
		ch <- struct{}{}
	}

	for _, dish := range dishes {
		wg.Add(1)
		go func(dish model.Dish) {
			defer wg.Done()

			<-ch

			err := repo.CreateOrUpdateDish(&dish)
			if err != nil {
			}

			ch <- struct{}{}
		}(dish)
	}

	wg.Wait()
	return nil
}
