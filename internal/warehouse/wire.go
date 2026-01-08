package warehouse

import (
	"github.com/Wenell09/MyStock/internal/warehouse/controller"
	"github.com/Wenell09/MyStock/internal/warehouse/repository"
	"github.com/Wenell09/MyStock/internal/warehouse/service"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	repository.NewWarehouseRepository,
	service.NewWarehouseService,
	controller.NewWarehouseController,
)
