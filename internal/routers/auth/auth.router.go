package register

import (
	"github.com/ducnv194023/shoppe_be_go/internal/wire"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (pr *AuthRouter) InitAuthRouter(router *gin.RouterGroup) {
	authController, _ := wire.InitializeAuthHandler()
	// public
	AuthRouterPublic := router.Group("/user")
	{
		AuthRouterPublic.GET("/register", authController.Register)
		AuthRouterPublic.GET("/sendOTP")
	}

	// private
	AuthRouterPrivate := router.Group("/user")
	// limit
	// UserRouterPrivate.Use(middleware.JwtAuthMiddleware())
	// permission

	{
		AuthRouterPrivate.GET("/profile")
	}
}
