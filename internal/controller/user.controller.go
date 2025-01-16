package controller

import (
	"github.com/ducnv194023/shoppe_be_go/internal/services"
	"github.com/gin-gonic/gin"
	response "github.com/ducnv194023/shoppe_be_go/pkg/response"
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

