package auth

import (
	"github.com/ducnv194023/shoppe_be_go/internal/wire"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (pr *AuthRouter) InitAuthRouter(router *gin.RouterGroup) {
	authController, _ := wire.InitializeAuthHandler()
	// public
	AuthRouterPublic := router.Group("/auth")
	{
		AuthRouterPublic.GET("/register", authController.Register)
		AuthRouterPublic.GET("/login", authController.Login)
		AuthRouterPublic.GET("/logout", authController.Logout)
		AuthRouterPublic.GET("/sendOTP")
		AuthRouterPublic.GET("/verify")

	}
}
