package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type Product struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Price      float64
	CategoryID uint       `gorm:"not null"`
	Category   Category   `gorm:"foreignKey:CategoryID"`
}

var db *gorm.DB

func Migrate() {
	conn := "root:@tcp(localhost:3306)/rest-go?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	db.AutoMigrate(&Category{}, &Product{})
	log.Println("Database Migrated")
}
