package service

import (
	"errors"

	categoryRepository "github.com/Wenell09/MyStock/internal/category/repository"
	"github.com/Wenell09/MyStock/internal/item/dto"
	itemRepository "github.com/Wenell09/MyStock/internal/item/repository"
	"github.com/Wenell09/MyStock/internal/models"
	supplierRepository "github.com/Wenell09/MyStock/internal/supplier/repository"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemServiceImpl struct {
	ItemRepository     itemRepository.ItemRepository
	CategoryRepository categoryRepository.CategoryRepository
	SupplierRepository supplierRepository.SupplierRepository
	Validate           *validator.Validate
}

func NewItemService(itemRepository itemRepository.ItemRepository,
	categoryRepository categoryRepository.CategoryRepository,
	supplierRepository supplierRepository.SupplierRepository,
	validate *validator.Validate) ItemService {
	return &ItemServiceImpl{
		ItemRepository:     itemRepository,
		CategoryRepository: categoryRepository,
		SupplierRepository: supplierRepository,
		Validate:           validate,
	}
}

// Create implements [ItemService].
func (i *ItemServiceImpl) Create(request dto.ItemRequest) (models.Item, error) {
	if err := i.Validate.Struct(&request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.Item{}, utils.NewFieldError(ve)
		}
		return models.Item{}, utils.ValidationError{Msg: err.Error()}
	}
	responseCategory, err := i.CategoryRepository.ReadByPublicId(request.CategoryPublicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Item{}, utils.NotFoundError{Msg: "Category Not Found!"}
		}
		return models.Item{}, err
	}
	responseSupplier, err := i.SupplierRepository.ReadByPublicId(request.SupplierPublicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Item{}, utils.NotFoundError{Msg: "Supplier Not Found!"}
		}
		return models.Item{}, err
	}
	data := models.Item{
		PublicID:   uuid.New().String(),
		Name:       request.Name,
		CategoryID: responseCategory.ID,
		SupplierID: responseSupplier.ID,
	}
	response, err := i.ItemRepository.Create(data)
	if err != nil {
		return models.Item{}, err
	}
	return response, nil
}

// Delete implements [ItemService].
func (i *ItemServiceImpl) Delete(publicId string) error {
	if publicId == "" {
		return utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := i.ItemRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFoundError{Msg: "Item Not Found!"}
		}
		return err
	}
	if err := i.ItemRepository.Delete(response); err != nil {
		return err
	}
	return nil
}

// DeleteAll implements [ItemService].
func (i *ItemServiceImpl) DeleteAll() error {
	if err := i.ItemRepository.DeleteAll(); err != nil {
		return err
	}
	return nil
}

// Read implements [ItemService].
func (i *ItemServiceImpl) Read() ([]models.Item, error) {
	response, err := i.ItemRepository.Read()
	if err != nil {
		return []models.Item{}, err
	}
	return response, nil
}

// ReadByPublicId implements [ItemService].
func (i *ItemServiceImpl) ReadByPublicId(publicId string) (models.Item, error) {
	if publicId == "" {
		return models.Item{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := i.ItemRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Item{}, utils.NotFoundError{Msg: "Item Not Found!"}
		}
		return models.Item{}, err
	}
	return response, nil
}

// Update implements [ItemService].
func (i *ItemServiceImpl) Update(publicId string, request dto.ItemRequest) (models.Item, error) {
	if publicId == "" {
		return models.Item{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	if err := i.Validate.Struct(&request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.Item{}, utils.NewFieldError(ve)
		}
		return models.Item{}, utils.ValidationError{Msg: err.Error()}
	}
	response, err := i.ItemRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Item{}, utils.NotFoundError{Msg: "Item Not Found!"}
		}
		return models.Item{}, err
	}
	responseCategory, err := i.CategoryRepository.ReadByPublicId(request.CategoryPublicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Item{}, utils.NotFoundError{Msg: "Category Not Found!"}
		}
		return models.Item{}, err
	}
	responseSupplier, err := i.SupplierRepository.ReadByPublicId(request.SupplierPublicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Item{}, utils.NotFoundError{Msg: "Supplier Not Found!"}
		}
		return models.Item{}, err
	}
	response.Name = request.Name
	response.CategoryID = responseCategory.ID
	response.SupplierID = responseSupplier.ID

	response, err = i.ItemRepository.Update(response)
	if err != nil {
		return models.Item{}, err
	}
	return response, nil
}

var _ ItemService = (*ItemServiceImpl)(nil)
