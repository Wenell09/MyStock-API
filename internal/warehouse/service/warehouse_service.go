package service

import (
	"github.com/Wenell09/MyStock/internal/models"
	"github.com/Wenell09/MyStock/internal/warehouse/dto"
)

type WarehouseService interface {
	Create(request dto.WarehouseRequest) (models.Warehouse, error)
	Read() ([]models.Warehouse, error)
	ReadByPublicId(publicId string) (models.Warehouse, error)
	ReadByWarehousePublicId(publicId string) (models.Warehouse, error)
	Update(publicId string, request dto.WarehouseRequest) (models.Warehouse, error)
	Delete(publicId string) error
	DeleteAll() error
}
