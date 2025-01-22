package validation

import (
	baseValidator "github.com/ducnv194023/shoppe_be_go/internal/middleware"

	request "github.com/ducnv194023/shoppe_be_go/internal/vo"
)

func ValidateRegister(req request.RegisterRequest) error {
	return baseValidator.GetValidator().Struct(req)
}

func ValidateLogin(req request.LoginRequest) error {
	return baseValidator.GetValidator().Struct(req)
}
