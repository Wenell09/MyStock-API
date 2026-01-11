package repository

import "github.com/Wenell09/MyStock/internal/models"

type SupplierRepository interface {
	Create(supplier models.Supplier) (models.Supplier, error)
	Read() ([]models.Supplier, error)
	ReadByPublicId(publicId string) (models.Supplier, error)
	ReadBySupplierPublicId(publicId string) (models.Supplier, error)
	Update(supplier models.Supplier) (models.Supplier, error)
	Delete(supplier models.Supplier) error
	DeleteAll() error
}
