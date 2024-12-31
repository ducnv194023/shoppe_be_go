package controller

import (
	"shoppe_be/internal/services"
	"github.com/gin-gonic/gin"
	response "shoppe_be/pkg/response"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 20001, uc.userService.GetUserById())
}

