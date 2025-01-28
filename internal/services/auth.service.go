package services

import (
	"context"
	"errors"
	"time"

	"github.com/ducnv194023/shoppe_be_go/internal/constants"
	"github.com/ducnv194023/shoppe_be_go/internal/repo"
	"github.com/redis/go-redis/v9"
)

type IAuthService interface {
	Register(
		ctx context.Context, 
		email string,
		password string,
		fullname string,
	) (*redis.StatusCmd, error)
}

type AuthService struct {
	userRepo    repo.IUserRepo
	otpRepo     repo.IOTPRepo
	mailService IMailService
	redisClient	repo.IRedisRepo
}

func NewAuthService(
	userRepo repo.IUserRepo,
	otpRepo repo.IOTPRepo,
	mailService IMailService,
	redisClient	repo.IRedisRepo,
) IAuthService {
	return &AuthService{
		userRepo:    userRepo,
		otpRepo:     otpRepo,
		mailService: mailService,
		redisClient: redisClient,
	}
}

// Register implements IUserService.
func (as *AuthService) Register(
	ctx context.Context, 
	email string,
	password string,
	fullname string,
) (*redis.StatusCmd, error) {
	// find exist email
	existEmail, _ := as.userRepo.GetUserBasicByEmail(
		ctx, 
		email,
	)

	if existEmail != nil {
		return nil, errors.New("email is already exist")
	}
		// generate otp
	otp := as.otpRepo.GenerateOTP()

	redisBody := map[string]string{
		"otp":      otp,
		"email":    email,
		"password": password,
		"fullname": fullname,
	}

	redisCmd, error := as.redisClient.SetKey(ctx, constants.REGISTER_INFO_KEY, redisBody, time.Minute)

	return redisCmd, error
}


