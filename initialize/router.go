package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ducnv194023/shoppe_be_go/global"
	"github.com/ducnv194023/shoppe_be_go/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	s := global.Config.ServerSetting

	if s.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middleware

	r.Use() // logging
	r.Use()  // cross
	r.Use() // limit global

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		managerRouter.InitAdminRouter(MainGroup)
		managerRouter.InitUserRouter(MainGroup)
	}

	return r
}
