package manager

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// public
	AdminRouterPublic := router.Group("/admin")
	{
		AdminRouterPublic.GET("/login")
	}

	// private
	AdminRouterPrivate := router.Group("/admin")
	// limit
	// UserRouterPrivate.Use(middleware.JwtAuthMiddleware())
	// permission

	{
		AdminRouterPrivate.POST("/active")
	}
}
