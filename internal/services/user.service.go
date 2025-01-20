package services

import (
	"github.com/ducnv194023/shoppe_be_go/internal/repo"
)

type IUserService interface {
	Register(email string, password string) int
}

type UserService struct {
	userRepo repo.IUserRepo
}

func NewUserService(
	userRepo repo.IUserRepo,
) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *UserService) Register(email string, password string) int {
	return us.userRepo.Register(email, password)
}

