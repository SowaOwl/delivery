package main

import (
	"delivery/internal/model"
	"delivery/internal/provider"
	"delivery/internal/service"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:33090)/delivery?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
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
		&model.User{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully migrated DB")

	dishProvider := provider.NewFileDishProvider("dishes.json")

	err = service.UpdateDishes(dishProvider, db)

	if err != nil {
		log.Fatal(err)
	}
}
