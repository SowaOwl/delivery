package main

import (
	"delivery/models"
	"delivery/providers"
	"delivery/services"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/delivery?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Courier{},
		&models.Delivery{},
		&models.Dish{},
		&models.Order{},
		&models.Restaurant{},
		&models.User{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully migrated DB")

	provider := providers.NewFileDishProvider("dishes.json")

	err = services.UpdateDishes(provider, db)

	if err != nil {
		log.Fatal(err)
	}
}
