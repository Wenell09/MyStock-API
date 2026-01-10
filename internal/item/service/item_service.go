package service

import (
	"github.com/Wenell09/MyStock/internal/dto"
	"github.com/Wenell09/MyStock/internal/models"
)

type ItemService interface {
	Create(request dto.ItemRequest) (models.Item, error)
	Read() ([]models.Item, error)
	ReadByPublicId(publicId string) (models.Item, error)
	Update(publicId string, request dto.ItemRequest) (models.Item, error)
	Delete(publicId string) error
	DeleteAll() error
}
