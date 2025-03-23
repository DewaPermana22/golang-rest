package services

import (
	"REST_API_GO/models"
	"REST_API_GO/repositories"
)

type CategoryService struct {
	Repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService{
	return &CategoryService{Repo: repo}
}


func (s *CategoryService) GetAllCategory() ([]models.Category, error) {
	return s.Repo.GetAll()
}

func (s *CategoryService) CreateCategory(categories *models.Category) error {
	return s.Repo.Create(categories)
}