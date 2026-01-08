package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
)

type CategoryRepository interface {
	Create(category *models.Category) (models.Category, error)
	Read() ([]models.Category, error)
	ReadByPublicId(publicId string) (models.Category, error)
	Update(category *models.Category) (models.Category, error)
	Delete(category *models.Category) error
	DeleteAll() error
}
