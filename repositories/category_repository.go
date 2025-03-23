package repositories

import (
	"REST_API_GO/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	result := r.DB.Preload("Products").Find(&categories)
	return categories, result.Error
}

func (r *CategoryRepository) Create(categories *models.Category) error {
	return r.DB.Create(categories).Error
}
