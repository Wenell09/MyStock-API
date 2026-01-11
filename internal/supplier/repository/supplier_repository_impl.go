package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
	"gorm.io/gorm"
)

type SupplierRepositoryImpl struct {
	DB *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &SupplierRepositoryImpl{DB: db}
}

// Create implements [SupplierRepository].
func (s *SupplierRepositoryImpl) Create(supplier models.Supplier) (models.Supplier, error) {
	if err := s.DB.Create(&supplier).Error; err != nil {
		return models.Supplier{}, err
	}
	return supplier, nil
}

// Delete implements [SupplierRepository].
func (s *SupplierRepositoryImpl) Delete(supplier models.Supplier) error {
	if err := s.DB.Delete(&supplier, supplier.ID).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAll implements [SupplierRepository].
func (s *SupplierRepositoryImpl) DeleteAll() error {
	if err := s.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Supplier{}).Error; err != nil {
		return err
	}
	return nil
}

// Read implements [SupplierRepository].
func (s *SupplierRepositoryImpl) Read() ([]models.Supplier, error) {
	var data []models.Supplier
	if err := s.DB.Model(&models.Supplier{}).Find(&data).Error; err != nil {
		return []models.Supplier{}, err
	}
	return data, nil
}

// ReadByPublicId implements [SupplierRepository].
func (s *SupplierRepositoryImpl) ReadByPublicId(publicId string) (models.Supplier, error) {
	var data models.Supplier
	if err := s.DB.Where("public_id = ?", publicId).First(&data).Error; err != nil {
		return models.Supplier{}, err
	}
	return data, nil
}

// ReadBySupplierPublicId implements [SupplierRepository].
func (s *SupplierRepositoryImpl) ReadBySupplierPublicId(publicId string) (models.Supplier, error) {
	var data models.Supplier
	if err := s.DB.Preload("Items").Preload("Items.Category").Where("suppliers.public_id = ?", publicId).First(&data).Error; err != nil {
		return models.Supplier{}, err
	}
	return data, nil
}

// Update implements [SupplierRepository].
func (s *SupplierRepositoryImpl) Update(supplier models.Supplier) (models.Supplier, error) {
	if err := s.DB.Where("id = ?", supplier.ID).Updates(&supplier).Error; err != nil {
		return models.Supplier{}, err
	}
	return supplier, nil
}

var _ SupplierRepository = (*SupplierRepositoryImpl)(nil)
