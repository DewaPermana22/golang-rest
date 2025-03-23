package utils

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "root:@tcp(localhost:3306)/rest-go?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	log.Println("Connected to database")
}

