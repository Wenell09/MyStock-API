package service

import (
	"github.com/Wenell09/MyStock/internal/models"
	"github.com/Wenell09/MyStock/internal/transaction/dto"
)

type TransactionService interface {
	CreateOrUpdate(request dto.TransactionRequest) (models.StockTransaction, error)
	Read() ([]models.StockTransaction, error)
	ReadByPublicId(publicId string) (models.StockTransaction, error)
}
