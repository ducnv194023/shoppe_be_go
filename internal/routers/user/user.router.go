package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public
	UserRouterPublic := router.Group("/user")
	{
		UserRouterPublic.GET("/register")
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
