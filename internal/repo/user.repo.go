package repo

type IUserRepo interface {
	GetUserBasicByEmail()
	CreateUserBasic()
	CreateUserProfile()
}

type UserRepo struct{}

func (ur *UserRepo) CreateUserBasic() {
	panic("unimplemented")
}

func (ur *UserRepo) CreateUserProfile() {
	panic("unimplemented")
}

func (ur *UserRepo) GetUserBasicByEmail() {
	panic("unimplemented")
}

func NewUserRepo() IUserRepo {
	return &UserRepo{}
}
