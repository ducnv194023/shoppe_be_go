package controller

import (
	validator "github.com/ducnv194023/shoppe_be_go/internal/middleware"
	"github.com/ducnv194023/shoppe_be_go/internal/services"
	"github.com/ducnv194023/shoppe_be_go/internal/vo"
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
        response.RespondWithError(c, response.ErrCodeBadRequest, errorMsg)

		return
    }
	
	_, error := ac.authService.Register(c , req.Email, req.Password, req.FullName )

	if error != nil {
		response.ErrorResponse(c, response.ErrCodeInternalServoce, "đã có lỗi xảy ra")
	}

	response.SuccessResponse(c, 200, "thành công", nil)
}

func (ac *AuthController) Login(c *gin.Context) {

}

func (ac *AuthController) Logout(c *gin.Context) {

}
