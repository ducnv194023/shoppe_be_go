package controller

import (
	"github.com/ducnv194023/shoppe_be_go/internal/services"
	"github.com/gin-gonic/gin"
	response "github.com/ducnv194023/shoppe_be_go/pkg/response"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(
	userService services.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("ducnv@coolmate.me", "123456")
	response.SuccessResponse(c, result, nil)
}

