package services

import (
	"github.com/ducnv194023/shoppe_be_go/internal/repo"
)

type IUserService interface {
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


