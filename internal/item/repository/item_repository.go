package repository

import "github.com/Wenell09/MyStock/internal/models"

type ItemRepository interface {
	Create(item models.Item) (models.Item, error)
	Read() ([]models.Item, error)
	ReadByPublicId(publicId string) (models.Item, error)
	Update(item models.Item) (models.Item, error)
	Delete(item models.Item) error
	DeleteAll() error
}
