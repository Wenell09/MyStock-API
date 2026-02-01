package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *gorm.DB, data models.StockTransaction) error
	ReadAll() ([]models.StockTransaction, error)
	ReadByPublicID(publicID string) (models.StockTransaction, error)
}
