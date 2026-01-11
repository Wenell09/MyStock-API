package controller

import (
	"github.com/Wenell09/MyStock/internal/category/service"
	"github.com/Wenell09/MyStock/internal/dto"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
	Logger          *logrus.Logger
}

func NewCategoryController(categoryService service.CategoryService, logger *logrus.Logger) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
		Logger:          logger,
	}
}

// Create implements [CategoryController].
func (c *CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	var request dto.CategoryRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{
			Msg: "Invalid request body",
		})
	}
	response, err := c.CategoryService.Create(request)
	data := dto.CategoryResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	if err != nil {
		c.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	c.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusCreated,
		"request":  request.Name,
		"response": response,
	}).Info("Success Create New Category")
	return ctx.Status(fiber.StatusCreated).JSON(
		utils.NewResponseSuccess(fiber.StatusCreated, "Success Create New Category!", data),
	)
}

// Delete implements [CategoryController].
func (c *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	err := c.CategoryService.Delete(publicId)
	if err != nil {
		c.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	c.Logger.WithFields(logrus.Fields{
		"status":   200,
		"request":  publicId,
		"response": "Success Delete Category",
	}).Info("Success Delete Category")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Delete Category", nil),
	)
}

func (c *CategoryControllerImpl) DeleteAll(ctx *fiber.Ctx) error {
	if err := c.CategoryService.DeleteAll(); err != nil {
		c.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	c.Logger.WithFields(logrus.Fields{
		"status":   200,
		"request":  nil,
		"response": "Success Delete All Category!",
	}).Info("Success Delete All Category!")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Delete All Category!", nil),
	)
}

// Read implements [CategoryController].
func (c *CategoryControllerImpl) Read(ctx *fiber.Ctx) error {
	response, err := c.CategoryService.Read()
	var data []dto.CategoryResponse
	for _, dataResponse := range response {
		data = append(data, dto.CategoryResponse{
			PublicId:  dataResponse.PublicID,
			Name:      dataResponse.Name,
			CreatedAt: dataResponse.CreatedAt,
		})
	}
	if err != nil {
		c.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	c.Logger.WithFields(logrus.Fields{
		"status":   200,
		"request":  nil,
		"response": response,
	}).Info("Success Get Category!")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Get Category!", data),
	)
}

// ReadItemByCategoryPublicId implements [CategoryController].
func (c *CategoryControllerImpl) ReadByCategoryPublicID(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	response, err := c.CategoryService.ReadByCategoryPublicID(publicId)
	if err != nil {
		c.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	var items []dto.ItemResponse
	for _, responseItems := range response.Items {
		items = append(items, dto.ItemResponse{
			PublicId: responseItems.PublicID,
			Name:     responseItems.Name,
		})
	}
	data := dto.CategoryWithItemsResponse{
		PublicId: response.PublicID,
		Name:     response.Name,
		Items:    items,
	}
	c.Logger.WithFields(logrus.Fields{
		"status":   200,
		"request":  publicId,
		"response": response,
	}).Info("Success Get Items By Category!")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Get Items By Category!", data),
	)
}

// Update implements [CategoryController].
func (c *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	var request dto.CategoryRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{
			Msg: "Invalid request body",
		})
	}
	response, err := c.CategoryService.Update(publicId, request)
	data := dto.CategoryResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	if err != nil {
		c.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	c.Logger.WithFields(logrus.Fields{
		"status":   200,
		"request":  request,
		"response": response,
	}).Info("Success Update Category!")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Update Category!", data),
	)
}

var _ CategoryController = (*CategoryControllerImpl)(nil)
