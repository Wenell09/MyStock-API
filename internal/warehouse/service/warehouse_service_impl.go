package service

import (
	"errors"

	"github.com/Wenell09/MyStock/internal/dto"
	"github.com/Wenell09/MyStock/internal/models"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/Wenell09/MyStock/internal/warehouse/repository"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WarehouseServiceImpl struct {
	WarehouseRepository repository.WarehouseRepository
	Validate            *validator.Validate
}

func NewWarehouseService(warehouseRepository repository.WarehouseRepository, validate *validator.Validate) WarehouseService {
	return &WarehouseServiceImpl{
		WarehouseRepository: warehouseRepository,
		Validate:            validate,
	}
}

// Create implements [WarehouseService].
func (w *WarehouseServiceImpl) Create(request dto.WarehouseRequest) (models.Warehouse, error) {
	if err := w.Validate.Struct(request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.Warehouse{}, utils.NewFieldError(ve)
		}
		return models.Warehouse{}, utils.ValidationError{
			Msg: err.Error(),
		}
	}
	data := models.Warehouse{
		PublicID: uuid.New().String(),
		Name:     request.Name,
	}
	response, err := w.WarehouseRepository.Create(data)
	if err != nil {
		return models.Warehouse{}, err
	}
	return response, nil
}

// Delete implements [WarehouseService].
func (w *WarehouseServiceImpl) Delete(publicId string) error {
	if publicId == "" {
		return utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := w.WarehouseRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFoundError{Msg: "Warehouse Not Found!"}
		}
		return err
	}
	if err := w.WarehouseRepository.Delete(response); err != nil {
		return err
	}
	return nil
}

// DeleteAll implements [WarehouseService].
func (w *WarehouseServiceImpl) DeleteAll() error {
	if err := w.WarehouseRepository.DeleteAll(); err != nil {
		return err
	}
	return nil
}

// Read implements [WarehouseService].
func (w *WarehouseServiceImpl) Read() ([]models.Warehouse, error) {
	response, err := w.WarehouseRepository.Read()
	if err != nil {
		return []models.Warehouse{}, err
	}
	return response, nil
}

// ReadByPublicId implements [WarehouseService].
func (w *WarehouseServiceImpl) ReadByPublicId(publicId string) (models.Warehouse, error) {
	if publicId == "" {
		return models.Warehouse{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := w.WarehouseRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Warehouse{}, utils.NotFoundError{Msg: "Warehouse Not Found!"}
		}
		return models.Warehouse{}, err
	}
	return response, nil
}

// Update implements [WarehouseService].
func (w *WarehouseServiceImpl) Update(publicId string, request dto.WarehouseRequest) (models.Warehouse, error) {
	if publicId == "" {
		return models.Warehouse{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	if err := w.Validate.Struct(request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.Warehouse{}, utils.NewFieldError(ve)
		}
		return models.Warehouse{}, utils.ValidationError{
			Msg: err.Error(),
		}
	}
	response, err := w.WarehouseRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Warehouse{}, utils.NotFoundError{Msg: "Warehouse Not Found!"}
		}
		return models.Warehouse{}, err
	}
	response.Name = request.Name
	response, err = w.WarehouseRepository.Update(response)
	if err != nil {
		return models.Warehouse{}, err
	}
	return response, nil
}

var _ WarehouseService = (*WarehouseServiceImpl)(nil)
