package cmd

import (
	"delivery/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDataBase() *gorm.DB {
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

	return db
}
