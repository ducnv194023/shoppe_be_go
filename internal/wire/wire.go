//go:build wireinject

package wire

import (
	"github.com/ducnv194023/shoppe_be_go/internal/controller"
	"github.com/ducnv194023/shoppe_be_go/internal/repo"
	"github.com/ducnv194023/shoppe_be_go/internal/services"
	"github.com/google/wire"
)



func InitializeUserHandler() (*controller.UserController, error) {
    wire.Build(
        repo.NewUserRepo,
        services.NewUserService,
        controller.NewUserController,
    )
    return new(controller.UserController), nil
}