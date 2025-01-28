    //go:build wireinject

    package wire

    import (
        "database/sql"
	    "github.com/redis/go-redis/v9"

        "github.com/ducnv194023/shoppe_be_go/internal/controller"
        "github.com/ducnv194023/shoppe_be_go/internal/database/sqlc"
        "github.com/ducnv194023/shoppe_be_go/internal/repo"
        "github.com/ducnv194023/shoppe_be_go/internal/services"
        "github.com/ducnv194023/shoppe_be_go/global"

        "github.com/google/wire"
    )

    func provideDB() (*sql.DB, error) {
        return global.MDb, nil
    }

    func provideRedis() (*redis.Client, error) {
        return global.Rdb, nil
    }

    func InitializeAuthHandler() (*controller.AuthController, error) {
        wire.Build(
            provideDB,
            provideRedis,
            wire.Bind(new(database.DBTX), new(*sql.DB)),
            database.New,
            repo.NewRedisRepo,
            repo.NewUserRepo,
            repo.NewOTPRepo,
            wire.Bind(new(repo.IUserRepo), new(*repo.UserRepo)),
            services.NewMailService,
            services.NewAuthService,
            controller.NewAuthController,
        )
        return &controller.AuthController{}, nil
    }