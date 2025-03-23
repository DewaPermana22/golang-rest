package services

import (
	"REST_API_GO/models"
	"REST_API_GO/repositories"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService{
	return &ProductService{Repo: repo}
}


func (s *ProductService) GetAllProducts(page, limit int) ([]models.ProductResponse, int, int, int64, error) {
	offset := (page - 1) * limit

	products, err := s.Repo.GetAll(limit, offset)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	total := int64(len(products))
	totalPages := (total + int64(limit) - 1) / int64(limit)

	return products, page, limit, totalPages, nil
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.Repo.Create(product)
}