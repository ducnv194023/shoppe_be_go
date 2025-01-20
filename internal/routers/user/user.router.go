package user

import (
	"github.com/ducnv194023/shoppe_be_go/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userController, _ := wire.InitializeUserHandler()
	// public
	UserRouterPublic := router.Group("/user")
	{
		UserRouterPublic.GET("/register", userController.Register)
		UserRouterPublic.GET("/sendOTP")
	}

	// private
	UserRouterPrivate := router.Group("/user")
	// limit
	// UserRouterPrivate.Use(middleware.JwtAuthMiddleware())
	// permission

	{
		UserRouterPrivate.GET("/profile")
	}
}
