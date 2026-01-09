package controller

import (
	"github.com/Wenell09/MyStock/internal/dto"
	"github.com/Wenell09/MyStock/internal/supplier/service"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SupplierControllerImpl struct {
	SupplierService service.SupplierService
	Logger          *logrus.Logger
}

func NewSupplierController(supplierService service.SupplierService, logger *logrus.Logger) SupplierController {
	return &SupplierControllerImpl{
		SupplierService: supplierService,
		Logger:          logger,
	}
}

// Create implements [SupplierController].
func (s *SupplierControllerImpl) Create(ctx *fiber.Ctx) error {
	var request dto.SupplierRequest
	if err := ctx.BodyParser(&request); err != nil {
		s.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{Msg: "Error Parser Body!"})
	}
	response, err := s.SupplierService.Create(request)
	if err != nil {
		s.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	data := dto.SupplierResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	s.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusCreated,
		"request":  request.Name,
		"response": response,
	}).Info("Success Create New Supplier!")
	return ctx.Status(fiber.StatusCreated).JSON(utils.NewResponseSuccess(fiber.StatusCreated, "Success Create New Supplier!", data))
}

// Delete implements [SupplierController].
func (s *SupplierControllerImpl) Delete(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	if err := s.SupplierService.Delete(publicId); err != nil {
		s.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	s.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": "Success Delete Supplier!",
	}).Info("Success Delete Supplier!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Delete Supplier!", nil))
}

// DeleteAll implements [SupplierController].
func (s *SupplierControllerImpl) DeleteAll(ctx *fiber.Ctx) error {
	if err := s.SupplierService.DeleteAll(); err != nil {
		s.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	s.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": "Success Delete All Supplier!",
	}).Info("Success Delete All Supplier!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Delete All Supplier!", nil))
}

// Read implements [SupplierController].
func (s *SupplierControllerImpl) Read(ctx *fiber.Ctx) error {
	var data []dto.SupplierResponse
	response, err := s.SupplierService.Read()
	if err != nil {
		s.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	for _, responseData := range response {
		data = append(data, dto.SupplierResponse{
			PublicId:  responseData.PublicID,
			Name:      responseData.Name,
			CreatedAt: responseData.CreatedAt,
		})
	}
	s.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": response,
	}).Info("Success Get Supplier!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Get Supplier!", data))
}

// ReadByPublicId implements [SupplierController].
func (s *SupplierControllerImpl) ReadByPublicId(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	response, err := s.SupplierService.ReadByPublicId(publicId)
	if err != nil {
		s.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	data := dto.SupplierResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	s.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  publicId,
		"response": response,
	}).Info("Success Get Detail Supplier!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Get Detail Supplier!", data))
}

// Update implements [SupplierController].
func (s *SupplierControllerImpl) Update(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	var request dto.SupplierRequest
	if err := ctx.BodyParser(&request); err != nil {
		s.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{Msg: "Error Parser Body!"})
	}
	response, err := s.SupplierService.Update(publicId, request)
	if err != nil {
		s.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	data := dto.SupplierResponse{
		PublicId:  response.PublicID,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
	}
	s.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  request,
		"response": response,
	}).Info("Success Update Supplier!")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Update Supplier!", data))
}

var _ SupplierController = (*SupplierControllerImpl)(nil)
