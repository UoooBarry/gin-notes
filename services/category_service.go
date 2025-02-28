package services

import (
	"uooobarry/gin-notes/repositories"
)

type CategoryService struct {
	CategoryRepo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: repo}
}
