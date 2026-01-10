package item

import (
	"github.com/Wenell09/MyStock/internal/item/controller"
	"github.com/Wenell09/MyStock/internal/item/repository"
	"github.com/Wenell09/MyStock/internal/item/service"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	controller.NewItemController,
	service.NewItemService,
	repository.NewItemRepository,
)
