package services

import (
	"github.com/ducnv194023/shoppe_be_go/internal/repo"
)
type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserById() string {
	return us.userRepo.GetUserById()
}
