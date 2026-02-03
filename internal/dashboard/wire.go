package dashboard

import (
	"github.com/Wenell09/MyStock/internal/dashboard/controller"
	"github.com/Wenell09/MyStock/internal/dashboard/repository"
	"github.com/Wenell09/MyStock/internal/dashboard/service"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	repository.NewDashboardRepository,
	service.NewDashboardService,
	controller.NewDashboardController,
)
