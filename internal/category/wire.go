package category

import (
	"github.com/Wenell09/MyStock/internal/category/controller"
	"github.com/Wenell09/MyStock/internal/category/repository"
	"github.com/Wenell09/MyStock/internal/category/service"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)
