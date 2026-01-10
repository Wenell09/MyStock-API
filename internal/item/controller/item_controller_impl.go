package controller

import (
	"github.com/Wenell09/MyStock/internal/dto"
	"github.com/Wenell09/MyStock/internal/item/service"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ItemControllerImpl struct {
	ItemService service.ItemService
	Logger      *logrus.Logger
}

// Create implements [ItemController].
func (i *ItemControllerImpl) Create(ctx *fiber.Ctx) error {
	var request dto.ItemRequest
	if err := ctx.BodyParser(&request); err != nil {
		i.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{Msg: "Error Body Parser!"})
	}
	response, err := i.ItemService.Create(request)
	if err != nil {
		i.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	data := dto.ItemResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	i.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusCreated,
		"request":  request.Name,
		"response": response,
	}).Info("Success Create New Item!")
	return ctx.Status(fiber.StatusCreated).JSON(utils.NewResponseSuccess(fiber.StatusCreated, "Success Create New Item!", data))
}

// Delete implements [ItemController].
func (i *ItemControllerImpl) Delete(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	if err := i.ItemService.Delete(publicId); err != nil {
		i.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	i.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  publicId,
		"response": "Success Delete Item!",
	}).Info("Success Delete Item!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Delete Item!", nil))
}

// DeleteAll implements [ItemController].
func (i *ItemControllerImpl) DeleteAll(ctx *fiber.Ctx) error {
	if err := i.ItemService.DeleteAll(); err != nil {
		i.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	i.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": "Success Delete All Item!",
	}).Info("Success Delete All Item!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Delete All Item!", nil))
}

// Read implements [ItemController].
func (i *ItemControllerImpl) Read(ctx *fiber.Ctx) error {
	var data []dto.ItemResponse
	response, err := i.ItemService.Read()
	if err != nil {
		i.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	for _, responseData := range response {
		var warehouses []dto.WarehouseResponse
		for _, responseDataWarehouse := range responseData.ItemWarehouses {
			warehouses = append(warehouses, dto.WarehouseResponse{
				PublicId: responseDataWarehouse.Warehouse.PublicID,
				Name:     responseDataWarehouse.Warehouse.Name,
			})
		}
		data = append(data, dto.ItemResponse{
			PublicId: responseData.PublicID,
			Name:     responseData.Name,
			CategoryResponse: dto.CategoryResponse{
				PublicId: responseData.Category.PublicID,
				Name:     responseData.Category.Name,
			},
			SupplierResponse: dto.SupplierResponse{
				PublicId: responseData.Supplier.PublicID,
				Name:     responseData.Supplier.Name,
			},
			WarehouseResponse: warehouses,
			CreatedAt:         responseData.CreatedAt,
		})
	}
	i.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": "Success Get Items!",
	}).Info("Success Get Items!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Get Items!", data))
}

// ReadByPublicId implements [ItemController].
func (i *ItemControllerImpl) ReadByPublicId(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	if publicId == "" {
		return utils.NewHandleError(ctx, utils.ValidationError{Msg: "public_id must be filled!"})
	}
	response, err := i.ItemService.ReadByPublicId(publicId)
	if err != nil {
		i.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	var warehouses []dto.WarehouseResponse
	for _, responseDataWarehouse := range response.ItemWarehouses {
		warehouses = append(warehouses, dto.WarehouseResponse{
			PublicId: responseDataWarehouse.Warehouse.PublicID,
			Name:     responseDataWarehouse.Warehouse.Name,
		})
	}
	data := dto.ItemResponse{
		PublicId: response.PublicID,
		Name:     response.Name,
		CategoryResponse: dto.CategoryResponse{
			PublicId: response.Category.PublicID,
			Name:     response.Category.Name,
		},
		SupplierResponse: dto.SupplierResponse{
			PublicId: response.Supplier.PublicID,
			Name:     response.Supplier.Name,
		},
		WarehouseResponse: warehouses,
		CreatedAt:         response.CreatedAt,
	}
	i.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": "Success Get Detail Item!",
	}).Info("Success Get Detail Item!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Get Detail Item!", data))
}

// Update implements [ItemController].
func (i *ItemControllerImpl) Update(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	if publicId == "" {
		return utils.NewHandleError(ctx, utils.ValidationError{Msg: "public_id must be filled!"})
	}
	var request dto.ItemRequest
	if err := ctx.BodyParser(&request); err != nil {
		i.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{Msg: "Error Body Parser!"})
	}
	response, err := i.ItemService.Update(publicId, request)
	if err != nil {
		i.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	data := dto.ItemResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	i.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  request.Name,
		"response": response,
	}).Info("Success Update Item!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Update Item!", data))
}

var _ ItemController = (*ItemControllerImpl)(nil)
