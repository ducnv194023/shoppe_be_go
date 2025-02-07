package controller

import (
	validator "github.com/ducnv194023/shoppe_be_go/internal/middleware"
	"github.com/ducnv194023/shoppe_be_go/internal/services"
	"github.com/ducnv194023/shoppe_be_go/internal/vo"
	"github.com/ducnv194023/shoppe_be_go/pkg/errors"
	response "github.com/ducnv194023/shoppe_be_go/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(
	authService services.IAuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var req vo.RegisterRequest

    if err := c.ShouldBindJSON(&req); err != nil {
		errorMsg := validator.TranslateValidationErrors(err)
        response.ResponseError(c, errors.NewBadRequestError(errorMsg))

		return
    }
	
	_, error := ac.authService.Register(c , req.Email, req.Password, req.FullName )

	if error != nil {
		response.ResponseError(c, errors.NewInternalError(error))
	}

	response.ResponseSuccess(c, 200, nil)
}

func (ac *AuthController) Login(c *gin.Context) {

}

func (ac *AuthController) Logout(c *gin.Context) {

}
