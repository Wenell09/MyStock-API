package service

import (
	"github.com/Wenell09/MyStock/internal/dto"
	"github.com/Wenell09/MyStock/internal/models"
)

type SupplierService interface {
	Create(request dto.SupplierRequest) (models.Supplier, error)
	Read() ([]models.Supplier, error)
	ReadByPublicId(publicId string) (models.Supplier, error)
	ReadBySupplierPublicId(publicId string) (models.Supplier, error)
	Update(publicId string, request dto.SupplierRequest) (models.Supplier, error)
	Delete(publicId string) error
	DeleteAll() error
}
