package controller

import (
	"github.com/ducnv194023/shoppe_be_go/internal/services"
	"github.com/gin-gonic/gin"
	response "github.com/ducnv194023/shoppe_be_go/pkg/response"
)

type RegisterRequest struct {
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=8"`
    FullName  string `json:"full_name" binding:"required"`
}

type RegisterResponse struct {
    Message string `json:"message"`
}

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
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, err.Error(), nil)
		return
	}
	
	err := ac.authService.Register(req.Email, req.FullName, req.Password)
	if err != nil {
		response.ErrorResponse(c, 400) // Changed to pass status code instead of error message and nil
		return
	}

	response.SuccessResponse(c, RegisterResponse{Message: "Register successfully"}, nil)
}
