package services

import (
	"shoppe_be/internal/repo"
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
