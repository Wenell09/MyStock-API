package service

import (
	"github.com/Wenell09/MyStock/internal/category/dto"
	"github.com/Wenell09/MyStock/internal/models"
)

type CategoryService interface {
	Create(request dto.CategoryRequest) (models.Category, error)
	Read() ([]models.Category, error)
	ReadByPublicID(publicId string) (models.Category, error)
	ReadByCategoryPublicID(publicId string) (models.Category, error)
	Update(publicId string, request dto.CategoryRequest) (models.Category, error)
	Delete(publicId string) error
	DeleteAll() error
}
