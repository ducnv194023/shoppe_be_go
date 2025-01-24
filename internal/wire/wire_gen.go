// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"database/sql"
	"github.com/ducnv194023/shoppe_be_go/global"
	"github.com/ducnv194023/shoppe_be_go/internal/controller"
	"github.com/ducnv194023/shoppe_be_go/internal/database/sqlc"
	"github.com/ducnv194023/shoppe_be_go/internal/repo"
	"github.com/ducnv194023/shoppe_be_go/internal/services"
	"github.com/redis/go-redis/v9"
)

// Injectors from auth.wire.go:

func InitializeAuthHandler() (*controller.AuthController, error) {
	db, err := provideDB()
	if err != nil {
		return nil, err
	}
	queries := database.New(db)
	userRepo := repo.NewUserRepo(queries)
	iotpRepo := repo.NewOTPRepo()
	iMailService := services.NewMailService()
	client, err := provideRedis()
	if err != nil {
		return nil, err
	}
	iAuthService := services.NewAuthService(userRepo, iotpRepo, iMailService, client)
	authController := controller.NewAuthController(iAuthService)
	return authController, nil
}

// auth.wire.go:

func provideDB() (*sql.DB, error) {
	return global.MDb, nil
}

func provideRedis() (*redis.Client, error) {
	return global.Rdb, nil
}
