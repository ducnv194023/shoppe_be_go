package repo

import (
    "math/rand"
    "time"
)

type IOTPRepo interface {
	GenerateOTP() string
}

type OTPRepo struct {}


func (ur *OTPRepo) GenerateOTP() string {
    // Generate 6-digit OTP
    const length = 6
    rand.Seed(time.Now().UnixNano())
    
    // Create numeric OTP
    digits := make([]rune, length)
    for i := range digits {
        digits[i] = rune('0' + rand.Intn(10))
    }
    
    return string(digits)
}
// khởi tạo
func NewOTPRepo() IOTPRepo {
	return &OTPRepo{}
}
