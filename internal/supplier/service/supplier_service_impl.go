package service

import (
	"errors"

	"github.com/Wenell09/MyStock/internal/dto"
	"github.com/Wenell09/MyStock/internal/models"
	"github.com/Wenell09/MyStock/internal/supplier/repository"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupplierServiceImpl struct {
	SupplierRepository repository.SupplierRepository
	Validate           *validator.Validate
}

func NewSupplierService(supplierRepository repository.SupplierRepository,
	validate *validator.Validate) SupplierService {
	return &SupplierServiceImpl{
		SupplierRepository: supplierRepository,
		Validate:           validate,
	}
}

// Create implements [SupplierService].
func (s *SupplierServiceImpl) Create(request dto.SupplierRequest) (models.Supplier, error) {
	if err := s.Validate.Struct(request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.Supplier{}, utils.NewFieldError(ve)
		}
		return models.Supplier{}, utils.ValidationError{Msg: err.Error()}
	}
	data := models.Supplier{
		PublicID: uuid.New().String(),
		Name:     request.Name,
	}
	response, err := s.SupplierRepository.Create(data)
	if err != nil {
		return models.Supplier{}, err
	}
	return response, nil
}

// Delete implements [SupplierService].
func (s *SupplierServiceImpl) Delete(publicId string) error {
	if publicId == "" {
		return utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := s.SupplierRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFoundError{Msg: "Supplier Not Found!"}
		}
		return err
	}
	if err := s.SupplierRepository.Delete(response); err != nil {
		return err
	}
	return nil
}

// DeleteAll implements [SupplierService].
func (s *SupplierServiceImpl) DeleteAll() error {
	if err := s.SupplierRepository.DeleteAll(); err != nil {
		return err
	}
	return nil
}

// Read implements [SupplierService].
func (s *SupplierServiceImpl) Read() ([]models.Supplier, error) {
	response, err := s.SupplierRepository.Read()
	if err != nil {
		return []models.Supplier{}, err
	}
	return response, nil
}

// ReadByPublicId implements [SupplierService].
func (s *SupplierServiceImpl) ReadByPublicId(publicId string) (models.Supplier, error) {
	if publicId == "" {
		return models.Supplier{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := s.SupplierRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Supplier{}, utils.NotFoundError{Msg: "Supplier Not Found!"}
		}
		return models.Supplier{}, err
	}
	return response, nil
}

// ReadBySupplierPublicId implements [SupplierService].
func (s *SupplierServiceImpl) ReadBySupplierPublicId(publicId string) (models.Supplier, error) {
	if publicId == "" {
		return models.Supplier{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := s.SupplierRepository.ReadBySupplierPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Supplier{}, utils.NotFoundError{Msg: "Supplier Not Found!"}
		}
		return models.Supplier{}, err
	}
	return response, nil
}

// Update implements [SupplierService].
func (s *SupplierServiceImpl) Update(publicId string, request dto.SupplierRequest) (models.Supplier, error) {
	if publicId == "" {
		return models.Supplier{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	if err := s.Validate.Struct(request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.Supplier{}, utils.NewFieldError(ve)
		}
		return models.Supplier{}, err
	}
	response, err := s.SupplierRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Supplier{}, utils.NotFoundError{Msg: "Supplier Not Found!"}
		}
		return models.Supplier{}, err
	}
	response.Name = request.Name
	response, err = s.SupplierRepository.Update(response)
	if err != nil {
		return models.Supplier{}, err
	}
	return response, nil
}

var _ SupplierService = (*SupplierServiceImpl)(nil)
