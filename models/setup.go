package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=127.0.0.1 dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err)
		panic("Failed to connect to database!")
	}

	//_ = db.AutoMigrate(&Patient{})
	//_ = db.AutoMigrate(&Doctor{})

	DB = db
}
