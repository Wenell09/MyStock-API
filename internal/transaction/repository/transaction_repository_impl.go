package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{DB: db}
}

// Create implements [TransactionRepository].
func (t *TransactionRepositoryImpl) Create(tx *gorm.DB, data models.StockTransaction) error {
	return tx.Create(&data).Error
}

// ReadAll implements [TransactionRepository].
func (t *TransactionRepositoryImpl) ReadAll() ([]models.StockTransaction, error) {
	var data []models.StockTransaction
	if err := t.DB.Preload("Item").Preload("Item.Category").Preload("Item.Supplier").Preload("Warehouse").Find(&data).Error; err != nil {
		return []models.StockTransaction{}, err
	}
	return data, nil
}

// ReadByPublicID implements [TransactionRepository].
func (t *TransactionRepositoryImpl) ReadByPublicID(publicID string) (models.StockTransaction, error) {
	var data models.StockTransaction
	if err := t.DB.Preload("Item").Preload("Item.Category").Preload("Item.Supplier").Preload("Warehouse").Where("public_id = ?", publicID).First(&data).Error; err != nil {
		return models.StockTransaction{}, err
	}
	return data, nil
}

var _ TransactionRepository = (*TransactionRepositoryImpl)(nil)
