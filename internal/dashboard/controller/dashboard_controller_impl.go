package controller

import (
	"github.com/Wenell09/MyStock/internal/dashboard/dto"
	"github.com/Wenell09/MyStock/internal/dashboard/service"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DashboardControllerImpl struct {
	DashboardService service.DashboardService
	Logger           *logrus.Logger
}

func NewDashboardController(dashboardService service.DashboardService,
	logger *logrus.Logger) DashboardController {
	return &DashboardControllerImpl{
		DashboardService: dashboardService,
		Logger:           logger,
	}
}

// CountData implements [DashboardController].
func (d *DashboardControllerImpl) CountData(ctx *fiber.Ctx) error {
	response, err := d.DashboardService.CountData()
	if err != nil {
		d.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	data := dto.DashboardResponse{
		TotalItems:        response["total_items"],
		TotalSuppliers:    response["total_suppliers"],
		TotalWarehouses:   response["total_warehouses"],
		TotalTransactions: response["total_transactions"],
	}
	d.Logger.WithFields(logrus.Fields{
		"status":   200,
		"request":  nil,
		"response": data,
	}).Info("Success Get Dashboard")
	return ctx.Status(fiber.StatusOK).JSON(
		utils.NewResponseSuccess(fiber.StatusOK, "Success Get Dashboard", data),
	)
}
