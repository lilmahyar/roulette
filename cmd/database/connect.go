package database

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"roulette/cmd/model"
)

var DbInstance *gorm.DB

func ConnectToDatabase() error {
	db, err := gorm.Open(sqlserver.Open(""))

	if err != nil {
		log.Println(err)
	}

	log.Println("Connected to database")

	DbInstance = db

	db.AutoMigrate(&model.Room{})
	return err
}
