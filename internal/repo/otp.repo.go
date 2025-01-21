package repo

type IOTPRepo interface {
	CreateOTP() string
}

type OTPRepo struct {}


func (ur *OTPRepo) CreateOTP() string {
	return "otp"
}

// khởi tạo
func NewOTPRepo() IOTPRepo {
	return &OTPRepo{}
}
