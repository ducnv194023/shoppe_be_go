package routers

import (
	"github.com/gin-gonic/gin"
	c "shoppe_be/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	
	v1:=r.Group("/v1")

	{
		v1.GET("/users/:id", c.NewUserController().GetUserById)
	}
	return r
}
