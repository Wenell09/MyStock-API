package supplier

import (
	"github.com/Wenell09/MyStock/internal/supplier/controller"
	"github.com/Wenell09/MyStock/internal/supplier/repository"
	"github.com/Wenell09/MyStock/internal/supplier/service"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	repository.NewSupplierRepository,
	service.NewSupplierService,
	controller.NewSupplierController,
)
