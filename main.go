package main

import (
	"delivery/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:33090)/delivery"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//err = db.AutoMigrate(
	//	&models.User{},
	//	&models.Courier{},
	//	&models.Delivery{},
	//	&models.Dish{},
	//	&models.Order{},
	//	&models.Restaurant{},
	//	&models.User{},
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Successfully migrated DB")
	//
	//provider := providers.NewFileDishProvider("dishes.json")
	//err = services.UpdateDishes(provider, db)
	//if err != nil {
	//	log.Fatal(err)
	//}

	dish := models.Dish{
		ID:          1,
		Name:        "TEST",
		Description: "TEST",
		Price:       100.12,
	}

	result = db.First(&dish)

	fmt.Println(err.Error())

	db.Create(&dish)
}
