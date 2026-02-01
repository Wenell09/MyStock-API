package service

import (
	"errors"

	itemRepository "github.com/Wenell09/MyStock/internal/item/repository"
	"github.com/Wenell09/MyStock/internal/models"
	"github.com/Wenell09/MyStock/internal/transaction/dto"
	transactionRepository "github.com/Wenell09/MyStock/internal/transaction/repository"
	"github.com/Wenell09/MyStock/internal/utils"
	warehouseRepository "github.com/Wenell09/MyStock/internal/warehouse/repository"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	DB                       *gorm.DB
	TransactionRepository    transactionRepository.TransactionRepository
	ItemRepository           itemRepository.ItemRepository
	ItemWarehousesRepository transactionRepository.ItemWarehouseRepository
	WarehouseRepository      warehouseRepository.WarehouseRepository
	Validate                 *validator.Validate
}

func NewTransactionService(DB *gorm.DB, transactionRepository transactionRepository.TransactionRepository,
	itemRepository itemRepository.ItemRepository,
	itemWarehousesRepository transactionRepository.ItemWarehouseRepository,
	warehouseRepository warehouseRepository.WarehouseRepository,
	validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		DB:                       DB,
		TransactionRepository:    transactionRepository,
		ItemRepository:           itemRepository,
		ItemWarehousesRepository: itemWarehousesRepository,
		WarehouseRepository:      warehouseRepository,
		Validate:                 validate,
	}
}

// CreateOrUpdate implements [TransactionService].
func (t *TransactionServiceImpl) CreateOrUpdate(request dto.TransactionRequest) (models.StockTransaction, error) {
	transaction := models.StockTransaction{}
	if err := t.Validate.Struct(request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.StockTransaction{}, utils.NewFieldError(ve)
		}
		return models.StockTransaction{}, utils.ValidationError{Msg: err.Error()}
	}
	item, err := t.ItemRepository.ReadByPublicId(request.ItemPublicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.StockTransaction{}, utils.ValidationError{Msg: "Item not found"}
		}
		return models.StockTransaction{}, err
	}
	warehouse, err := t.WarehouseRepository.ReadByPublicId(request.WarehousePublicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.StockTransaction{}, utils.ValidationError{Msg: "Warehouse not found"}
		}
		return models.StockTransaction{}, err
	}
	err = t.DB.Transaction(func(tx *gorm.DB) error {
		responseItemWarehouses, err := t.ItemWarehousesRepository.FindByItemAndWarehouse(
			tx,
			item.ID,
			warehouse.ID,
		)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if request.Type == string(models.StockOut) {
					return utils.ValidationError{Msg: "Stock not available"}
				}
				dataItemWarehouse := models.ItemWarehouse{
					PublicID:    uuid.New().String(),
					ItemID:      item.ID,
					WarehouseID: warehouse.ID,
					Stock:       request.Quantity,
				}
				if err := t.ItemWarehousesRepository.Create(tx, dataItemWarehouse); err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			if request.Type == string(models.StockOut) {
				if responseItemWarehouses.Stock < request.Quantity {
					return utils.ValidationError{Msg: "Stock not sufficient"}
				}
				responseItemWarehouses.Stock -= request.Quantity
			} else {
				responseItemWarehouses.Stock += request.Quantity
			}
			if err := t.ItemWarehousesRepository.Update(tx, responseItemWarehouses); err != nil {
				return err
			}
		}
		transaction = models.StockTransaction{
			PublicID:    uuid.New().String(),
			ItemID:      item.ID,
			WarehouseID: warehouse.ID,
			Type:        models.StockTransactionType(request.Type),
			Quantity:    request.Quantity,
		}
		return t.TransactionRepository.Create(tx, transaction)
	})
	if err != nil {
		return models.StockTransaction{}, err
	}
	return transaction, nil
}

// Read implements [TransactionService].
func (t *TransactionServiceImpl) Read() ([]models.StockTransaction, error) {
	response, err := t.TransactionRepository.ReadAll()
	if err != nil {
		return []models.StockTransaction{}, err
	}
	return response, nil
}

// ReadByPublicId implements [TransactionService].
func (t *TransactionServiceImpl) ReadByPublicId(publicId string) (models.StockTransaction, error) {
	if publicId == "" {
		return models.StockTransaction{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := t.TransactionRepository.ReadByPublicID(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.StockTransaction{}, utils.NotFoundError{Msg: "Transaction not found!"}
		}
		return models.StockTransaction{}, err
	}
	return response, nil
}

var _ TransactionService = (*TransactionServiceImpl)(nil)
