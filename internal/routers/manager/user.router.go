package manager

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// private
	UserRouterPrivate := router.Group("/admin/user")
	// limit
	// UserRouterPrivate.Use(middleware.JwtAuthMiddleware())
	// permission

	{
		UserRouterPrivate.POST("/active")
	}
}
