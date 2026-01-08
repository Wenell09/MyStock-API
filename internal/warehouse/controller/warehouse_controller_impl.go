package controller

import (
	"github.com/Wenell09/MyStock/internal/dto"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/Wenell09/MyStock/internal/warehouse/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WarehouseControllerImpl struct {
	WarehouseService service.WarehouseService
	Logger           *logrus.Logger
}

func NewWarehouseController(warehouseService service.WarehouseService,
	logger *logrus.Logger) WarehouseController {
	return &WarehouseControllerImpl{
		WarehouseService: warehouseService,
		Logger:           logger,
	}
}

// Create implements [WarehouseController].
func (w *WarehouseControllerImpl) Create(ctx *fiber.Ctx) error {
	var request dto.WarehouseRequest
	if err := ctx.BodyParser(&request); err != nil {
		w.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{
			Msg: "Invalid request body",
		})
	}
	response, err := w.WarehouseService.Create(request)
	data := dto.WarehouseResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	if err != nil {
		w.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	w.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusCreated,
		"request":  request.Name,
		"response": response,
	}).Info("Success Create Warehouse!")
	return ctx.Status(fiber.StatusCreated).JSON(
		utils.NewResponseSuccess(fiber.StatusCreated, "Success Create New Warehouse!", data),
	)
}

// Delete implements [WarehouseController].
func (w *WarehouseControllerImpl) Delete(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	if err := w.WarehouseService.Delete(publicId); err != nil {
		w.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	w.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": "Success Delete Warehouse!",
	}).Info("Success Delete Warehouse!")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Delete Warehouse!", nil),
	)
}

// DeleteAll implements [WarehouseController].
func (w *WarehouseControllerImpl) DeleteAll(ctx *fiber.Ctx) error {
	if err := w.WarehouseService.DeleteAll(); err != nil {
		w.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	w.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": "Success Delete All Warehouse!",
	}).Info("Success Delete All Warehouse!")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Delete All Warehouse!", nil),
	)
}

// Read implements [WarehouseController].
func (w *WarehouseControllerImpl) Read(ctx *fiber.Ctx) error {
	var data []dto.WarehouseResponse
	response, err := w.WarehouseService.Read()
	if err != nil {
		w.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	for _, dataResponse := range response {
		data = append(data, dto.WarehouseResponse{
			PublicId:  dataResponse.PublicID,
			Name:      dataResponse.Name,
			CreatedAt: dataResponse.CreatedAt,
		})
	}
	w.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": response,
	}).Info("Success Get Warehouse!")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Get Warehouse!", data),
	)
}

// ReadByPublicId implements [WarehouseController].
func (w *WarehouseControllerImpl) ReadByPublicId(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	response, err := w.WarehouseService.ReadByPublicId(publicId)
	data := dto.WarehouseResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	if err != nil {
		w.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	w.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": response,
	}).Info("Success Get Warehouse!")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Get Warehouse!", data),
	)
}

// Update implements [WarehouseController].
func (w *WarehouseControllerImpl) Update(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	var request dto.WarehouseRequest
	if err := ctx.BodyParser(&request); err != nil {
		w.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{
			Msg: "Invalid request body",
		})
	}
	response, err := w.WarehouseService.Update(publicId, request)
	data := dto.WarehouseResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	if err != nil {
		w.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	w.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": response,
	}).Info("Success Update Warehouse!")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Update Warehouse!", data),
	)
}

var _ WarehouseController = (*WarehouseControllerImpl)(nil)
