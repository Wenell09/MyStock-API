package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

// Create implements [CategoryRepository].
func (c *CategoryRepositoryImpl) Create(category *models.Category) (models.Category, error) {
	if err := c.db.Create(&category).Error; err != nil {
		return models.Category{}, err
	}
	return *category, nil
}

// Delete implements [CategoryRepository].
func (c *CategoryRepositoryImpl) Delete(category *models.Category) error {
	if err := c.db.Delete(&models.Category{}, category.ID).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAll implements [CategoryRepository].
func (c *CategoryRepositoryImpl) DeleteAll() error {
	if err := c.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Category{}).Error; err != nil {
		return err
	}
	return nil
}

// Read implements [CategoryRepository].
func (c *CategoryRepositoryImpl) Read() ([]models.Category, error) {
	var data []models.Category
	if err := c.db.Model(&models.Category{}).Find(&data).Error; err != nil {
		return []models.Category{}, err
	}
	return data, nil
}

// ReadById implements [CategoryRepository].
func (c *CategoryRepositoryImpl) ReadByPublicId(publicId string) (models.Category, error) {
	var category models.Category
	if err := c.db.Where("public_id = ?", publicId).First(&category).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

// ReadItemByCategoryPublicId implements [CategoryRepository].
func (c *CategoryRepositoryImpl) ReadByCategoryPublicID(publicId string) (models.Category, error) {
	var category models.Category
	if err := c.db.Preload("Items").Where("categories.public_id = ?", publicId).First(&category).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

// Update implements [CategoryRepository].
func (c *CategoryRepositoryImpl) Update(category *models.Category) (models.Category, error) {
	if err := c.db.Where("id = ?", category.ID).Updates(&category).Error; err != nil {
		return models.Category{}, err
	}
	return *category, nil
}

var _ CategoryRepository = (*CategoryRepositoryImpl)(nil)
