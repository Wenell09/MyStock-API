package auth

import (
	"github.com/Wenell09/MyStock/internal/auth/controller"
	"github.com/Wenell09/MyStock/internal/auth/provider"
	"github.com/Wenell09/MyStock/internal/auth/repository"
	"github.com/Wenell09/MyStock/internal/auth/service"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	repository.NewAuthRepository,
	provider.NewAuthProvider,
	service.NewAuthService,
	controller.NewAuthController,
)
