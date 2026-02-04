package repository

import "github.com/Wenell09/MyStock/internal/models"

type WarehouseRepository interface {
	Create(Warehouse models.Warehouse) (models.Warehouse, error)
	Read() ([]models.Warehouse, error)
	ReadByPublicId(publicId string) (models.Warehouse, error)
	ReadByWarehousePublicId(publicId string) (models.Warehouse, error)
	Update(Warehouse models.Warehouse) (models.Warehouse, error)
	Delete(Warehouse models.Warehouse) error
	DeleteAll() error
}
