package repo

import (
	utils "github.com/ducnv194023/shoppe_be_go/internal/utils/random"
)

type IOTPRepo interface {
	GenerateOTP() string
}

type OTPRepo struct {}

func (ur *OTPRepo) GenerateOTP() string {
    return utils.Random(6)
}
// khởi tạo
func NewOTPRepo() IOTPRepo {
	return &OTPRepo{}
}
