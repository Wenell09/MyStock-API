package service

import (
	"errors"

	"github.com/Wenell09/MyStock/internal/category/repository"
	"github.com/Wenell09/MyStock/internal/dto"
	"github.com/Wenell09/MyStock/internal/models"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		validate:           validate,
	}
}

// Create implements [CategoryService].
func (c *CategoryServiceImpl) Create(request dto.CategoryRequest) (models.Category, error) {
	data := models.Category{
		PublicID: uuid.New().String(),
		Name:     request.Name,
	}
	if err := c.validate.Struct(data); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.Category{}, utils.NewFieldError(ve)
		}
		return models.Category{}, utils.ValidationError{
			Msg: err.Error(),
		}
	}
	response, err := c.CategoryRepository.Create(&data)
	if err != nil {
		return models.Category{}, err
	}
	return response, nil
}

// Delete implements [CategoryService].
func (c *CategoryServiceImpl) Delete(publicId string) error {
	if publicId == "" {
		return utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := c.CategoryRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFoundError{Msg: "Category not found"}
		}
		return err
	}
	if err := c.CategoryRepository.Delete(&response); err != nil {
		return err
	}
	return nil
}

func (c *CategoryServiceImpl) DeleteAll() error {
	if err := c.CategoryRepository.DeleteAll(); err != nil {
		return err
	}
	return nil
}

// Read implements [CategoryService].
func (c *CategoryServiceImpl) Read() ([]models.Category, error) {
	response, err := c.CategoryRepository.Read()
	if err != nil {
		return []models.Category{}, err
	}
	return response, nil
}

// ReadByPublicID implements [CategoryService].
func (c *CategoryServiceImpl) ReadByPublicID(publicId string) (models.Category, error) {
	if publicId == "" {
		return models.Category{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	response, err := c.CategoryRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Category{}, utils.NotFoundError{Msg: "Category Not Found!"}
		}
		return models.Category{}, err
	}
	return response, nil
}

// Update implements [CategoryService].
func (c *CategoryServiceImpl) Update(publicId string, request dto.CategoryRequest) (models.Category, error) {
	if publicId == "" {
		return models.Category{}, utils.ValidationError{Msg: "public_id must be filled!"}
	}
	if err := c.validate.Struct(request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.Category{}, utils.NewFieldError(ve)
		}
		return models.Category{}, utils.ValidationError{
			Msg: err.Error(),
		}
	}
	response, err := c.CategoryRepository.ReadByPublicId(publicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Category{}, utils.NotFoundError{Msg: "Category Not Found!"}
		}
		return models.Category{}, err
	}
	response.Name = request.Name
	response, err = c.CategoryRepository.Update(&response)
	if err != nil {
		return models.Category{}, err
	}
	return response, nil
}

var _ CategoryService = (*CategoryServiceImpl)(nil)
