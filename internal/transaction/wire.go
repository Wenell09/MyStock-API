package transaction

import (
	"github.com/Wenell09/MyStock/internal/transaction/controller"
	"github.com/Wenell09/MyStock/internal/transaction/repository"
	"github.com/Wenell09/MyStock/internal/transaction/service"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	repository.NewTransactionRepository,
	repository.NewItemWarehouseRepository,
	service.NewTransactionService,
	controller.NewTransactionController,
)
