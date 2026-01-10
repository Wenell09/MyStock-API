package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
	"gorm.io/gorm"
)

type ItemRepositoryImpl struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &ItemRepositoryImpl{DB: db}
}

// Create implements [ItemRepository].
func (i *ItemRepositoryImpl) Create(item models.Item) (models.Item, error) {
	if err := i.DB.Create(&item).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

// Delete implements [ItemRepository].
func (i *ItemRepositoryImpl) Delete(item models.Item) error {
	if err := i.DB.Delete(&item, item.ID).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAll implements [ItemRepository].
func (i *ItemRepositoryImpl) DeleteAll() error {
	if err := i.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Item{}).Error; err != nil {
		return err
	}
	return nil
}

// Read implements [ItemRepository].
func (i *ItemRepositoryImpl) Read() ([]models.Item, error) {
	var data []models.Item
	if err := i.DB.Preload("Category").Preload("Supplier").Preload("ItemWarehouses").Preload("ItemWarehouses.Warehouse").Find(&data).Error; err != nil {
		return []models.Item{}, err
	}
	return data, nil
}

// ReadByPublicId implements [ItemRepository].
func (i *ItemRepositoryImpl) ReadByPublicId(publicId string) (models.Item, error) {
	var data models.Item
	if err := i.DB.Preload("Category").Preload("Supplier").Preload("ItemWarehouses").Preload("ItemWarehouses.Warehouse").Where("items.public_id = ?", publicId).Find(&data).Error; err != nil {
		return models.Item{}, err
	}
	return data, nil
}

// Update implements [ItemRepository].
func (i *ItemRepositoryImpl) Update(item models.Item) (models.Item, error) {
	if err := i.DB.Where("id = ?", item.ID).Updates(&item).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

var _ ItemRepository = (*ItemRepositoryImpl)(nil)
