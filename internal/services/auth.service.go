package services

import (
	"github.com/ducnv194023/shoppe_be_go/internal/repo"
)

type IAuthService interface {
	Register(email string, name string, password string)
	Login()
	Logout()
}

type AuthService struct {
	userRepo    repo.IUserRepo
	otpRepo     repo.IOTPRepo
	mailService IMailService
}

// Login implements IAuthService.
func (as *AuthService) Login() {
	panic("unimplemented")
}

// Logout implements IAuthService.
func (as *AuthService) Logout() {
	panic("unimplemented")
}

func NewAuthService(
	userRepo repo.IUserRepo,
	otpRepo repo.IOTPRepo,
	mailService IMailService,
) IAuthService {
	return &AuthService{
		userRepo:    userRepo,
		otpRepo:     otpRepo,
		mailService: mailService,
	}
}

// Register implements IUserService.
func (as *AuthService) Register(
	email string,
	name string,
	password string,
) {
	// validate input
	// find exist email
	as.userRepo.GetUserBasicByEmail()
	// generate otp
	as.otpRepo.CreateOTP()

	as.userRepo.CreateUserBasic()

	as.userRepo.CreateUserProfile()

	as.mailService.SendVerificationEmail()
}
