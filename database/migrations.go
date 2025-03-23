package database

import (
	"REST_API_GO/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MigrationDB() {
	conn := "root:@tcp(localhost:3306)/rest-go?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	// AutoMigrate untuk membuat tabel
	DB.AutoMigrate(&models.Category{}, &models.Product{})
	log.Println("Database Migrated")

	SeedDatabase()
}


func SeedDatabase() {
	var count int64
	DB.Model(&models.Category{}).Count(&count)
	if count > 0 {
		fmt.Println("ℹ️ Data kategori dan produk sudah ada, tidak perlu insert lagi.")
		return
	}

	// Insert 5 Kategori
	categories := []models.Category{
		{Name: "Makanan"},
		{Name: "Minuman"},
		{Name: "Snack"},
		{Name: "Buah"},
		{Name: "Sayur"},
	}
	DB.Create(&categories)
	fmt.Println("✅ 5 Kategori berhasil ditambahkan!")

	// Insert 15 Produk
	products := []models.Product{
		{Name: "Nasi Goreng", Price: 15000, CategoryID: 1},
		{Name: "Mie Goreng", Price: 12000, CategoryID: 1},
		{Name: "Ayam Bakar", Price: 20000, CategoryID: 1},
		{Name: "Jus Alpukat", Price: 10000, CategoryID: 2},
		{Name: "Teh Manis", Price: 5000, CategoryID: 2},
		{Name: "Kopi Hitam", Price: 8000, CategoryID: 2},
		{Name: "Keripik Kentang", Price: 7000, CategoryID: 3},
		{Name: "Popcorn", Price: 10000, CategoryID: 3},
		{Name: "Kacang Panggang", Price: 12000, CategoryID: 3},
		{Name: "Apel Merah", Price: 9000, CategoryID: 4},
		{Name: "Pisang Raja", Price: 8000, CategoryID: 4},
		{Name: "Semangka", Price: 15000, CategoryID: 4},
		{Name: "Bayam", Price: 5000, CategoryID: 5},
		{Name: "Wortel", Price: 6000, CategoryID: 5},
		{Name: "Brokoli", Price: 10000, CategoryID: 5},
	}
	DB.Create(&products)
	fmt.Println("✅ 15 Produk berhasil ditambahkan!")
}

