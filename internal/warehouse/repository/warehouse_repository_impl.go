package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
	"gorm.io/gorm"
)

type WarehouseRepositoryImpl struct {
	DB *gorm.DB
}

func NewWarehouseRepository(DB *gorm.DB) WarehouseRepository {
	return &WarehouseRepositoryImpl{
		DB: DB,
	}
}

// Create implements [WarehouseRepository].
func (w *WarehouseRepositoryImpl) Create(Warehouse models.Warehouse) (models.Warehouse, error) {
	if err := w.DB.Create(&Warehouse).Error; err != nil {
		return models.Warehouse{}, err
	}
	return Warehouse, nil
}

// Delete implements [WarehouseRepository].
func (w *WarehouseRepositoryImpl) Delete(Warehouse models.Warehouse) error {
	if err := w.DB.Delete(&Warehouse, Warehouse.ID).Error; err != nil {
		return err
	}
	return nil
}

func (w *WarehouseRepositoryImpl) DeleteAll() error {
	if err := w.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Warehouse{}).Error; err != nil {
		return err
	}
	return nil
}

// Read implements [WarehouseRepository].
func (w *WarehouseRepositoryImpl) Read() ([]models.Warehouse, error) {
	var data []models.Warehouse
	if err := w.DB.Model(&models.Warehouse{}).Find(&data).Error; err != nil {
		return []models.Warehouse{}, err
	}
	return data, nil
}

// ReadByPublicId implements [WarehouseRepository].
func (w *WarehouseRepositoryImpl) ReadByPublicId(publicId string) (models.Warehouse, error) {
	var data models.Warehouse
	if err := w.DB.Where("public_id = ?", publicId).First(&data).Error; err != nil {
		return models.Warehouse{}, err
	}
	return data, nil
}

// ReadByWarehousePublicId implements [WarehouseRepository].
func (w *WarehouseRepositoryImpl) ReadByWarehousePublicId(publicId string) (models.Warehouse, error) {
	var data models.Warehouse
	if err := w.DB.Preload("ItemWarehouses").Preload("ItemWarehouses.Item").Preload("ItemWarehouses.Item.Category").Preload("ItemWarehouses.Item.Supplier").Where("public_id = ?", publicId).First(&data).Error; err != nil {
		return models.Warehouse{}, err
	}
	return data, nil
}

// Update implements [WarehouseRepository].
func (w *WarehouseRepositoryImpl) Update(Warehouse models.Warehouse) (models.Warehouse, error) {
	if err := w.DB.Where("id = ?", Warehouse.ID).Updates(&Warehouse).Error; err != nil {
		return models.Warehouse{}, err
	}
	return Warehouse, nil
}

var _ WarehouseRepository = (*WarehouseRepositoryImpl)(nil)
