package main

import (
	"delivery/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:33090)/delivery?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Courier{},
		&model.Delivery{},
		&model.Dish{},
		&model.Order{},
		&model.Restaurant{},
	)
	if err != nil {
		log.Fatal(err)
	}

	//dishProvider := provider.NewFileDishProvider("dishes.json")
	//
	//err = service.UpdateDishes(dishProvider, db)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	//orderDTO := service.CreateOrderDTO{
	//	UserId:   1,
	//	Dishes:   []model.Dish{{ID: 42}, {ID: 929}, {ID: 248}},
	//	OrderSum: 76.2,
	//}
	//
	//err = service.CreateOrder(orderDTO, db)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//order := model.Order{ID: 1}
	//
	//err = service.CancelledOrder(order, db)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
