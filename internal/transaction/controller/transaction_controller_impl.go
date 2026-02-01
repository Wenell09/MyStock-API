package controller

import (
	"github.com/Wenell09/MyStock/internal/transaction/dto"
	"github.com/Wenell09/MyStock/internal/transaction/service"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
	Logger             *logrus.Logger
}

func NewTransactionController(transactionService service.TransactionService,
	logger *logrus.Logger) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
		Logger:             logger,
	}
}

// CreateOrUpdate implements [TransactionController].
func (t *TransactionControllerImpl) CreateOrUpdate(ctx *fiber.Ctx) error {
	var request dto.TransactionRequest
	if err := ctx.BodyParser(&request); err != nil {
		t.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{
			Msg: "Invalid request body",
		})
	}
	response, err := t.TransactionService.CreateOrUpdate(request)
	if err != nil {
		t.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	data := dto.TransactionCreatedOrUpdatedResponse{
		PublicId:          response.PublicID,
		ItemPublicId:      request.ItemPublicId,
		WarehousePublicId: request.WarehousePublicId,
		Type:              string(response.Type),
		Quantity:          response.Quantity,
		CreatedAt:         response.CreatedAt,
	}
	t.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusCreated,
		"request":  request,
		"response": data,
	}).Info("Success Create New Transaction")
	return ctx.Status(fiber.StatusCreated).JSON(
		utils.NewResponseSuccess(fiber.StatusCreated, "Success Create Transaction!", data),
	)
}

// Read implements [TransactionController].
func (t *TransactionControllerImpl) Read(ctx *fiber.Ctx) error {
	data := []dto.TransactionResponse{}
	response, err := t.TransactionService.Read()
	if err != nil {
		t.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	for _, dataResponse := range response {
		data = append(data, dto.TransactionResponse{
			PublicID: dataResponse.PublicID,
			ItemResponse: dto.ItemResponse{
				PublicId: dataResponse.Item.PublicID,
				Name:     dataResponse.Item.Name,
				CategoryResponse: dto.CategoryResponse{
					PublicId: dataResponse.Item.Category.PublicID,
					Name:     dataResponse.Item.Category.Name,
				},
				SupplierResponse: dto.SupplierResponse{
					PublicId: dataResponse.Item.Supplier.PublicID,
					Name:     dataResponse.Item.Supplier.Name,
				},
			},
			WarehouseResponse: dto.WarehouseResponse{
				PublicId: dataResponse.Warehouse.PublicID,
				Name:     dataResponse.Warehouse.Name,
			},
			Type:      string(dataResponse.Type),
			Quantity:  dataResponse.Quantity,
			CreatedAt: dataResponse.CreatedAt,
		})
	}
	t.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": data,
	}).Info("Success Get Transaction")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Get Transaction", data))
}

// ReadByPublicId implements [TransactionController].
func (t *TransactionControllerImpl) ReadByPublicId(ctx *fiber.Ctx) error {
	publicId := ctx.Params("public_id")
	response, err := t.TransactionService.ReadByPublicId(publicId)
	if err != nil {
		t.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	data := dto.TransactionResponse{
		PublicID: response.PublicID,
		ItemResponse: dto.ItemResponse{
			PublicId: response.Item.PublicID,
			Name:     response.Item.Name,
			CategoryResponse: dto.CategoryResponse{
				PublicId: response.Item.Category.PublicID,
				Name:     response.Item.Category.Name,
			},
			SupplierResponse: dto.SupplierResponse{
				PublicId: response.Item.Supplier.PublicID,
				Name:     response.Item.Supplier.Name,
			},
		},
		WarehouseResponse: dto.WarehouseResponse{
			PublicId: response.Warehouse.PublicID,
			Name:     response.Warehouse.Name,
		},
		Type:      string(response.Type),
		Quantity:  response.Quantity,
		CreatedAt: response.CreatedAt,
	}
	t.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusOK,
		"request":  nil,
		"response": data,
	}).Info("Success Get Detail Transaction")
	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponseSuccess(fiber.StatusOK, "Success Get Detail Transaction", data))
}
