//go:build wireinject

package wire

import (
	"github.com/ducnv194023/shoppe_be_go/internal/controller"
	"github.com/ducnv194023/shoppe_be_go/internal/repo"
	"github.com/ducnv194023/shoppe_be_go/internal/services"
	"github.com/google/wire"
)

func InitializeAuthHandler() (*controller.AuthController, error) {
    wire.Build(
        repo.NewUserRepo,
        repo.NewOTPRepo,
        services.NewMailService,
        services.NewAuthService,
        controller.NewAuthController,
    )
    return new(controller.AuthController), nil
}