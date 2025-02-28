package repositories

import (
	"uooobarry/gin-notes/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (repo *CategoryRepository) Create(category *models.Category) error {
	return repo.DB.Create(category).Error
}

func (repo *CategoryRepository) List() ([]models.Category, error) {
	var categories []models.Category

	err := repo.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
