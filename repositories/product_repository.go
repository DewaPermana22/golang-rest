package repositories

import (
	"REST_API_GO/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetAll(limit, offset int) ([]models.ProductResponse, error) {
	var products []models.ProductResponse
	var total int64

	//Hitung jumlah data
	r.DB.Table("products").Count(&total)

	result := r.DB.Table("products").
		Select("products.id, products.name, products.price, products.category_id, categories.name as category_name").
		Joins("left join categories on categories.id = products.category_id").
		Limit(limit).
		Offset(offset).
		Find(&products)
	return products, result.Error
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}
