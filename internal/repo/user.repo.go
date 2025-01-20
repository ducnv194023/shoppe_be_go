package repo

type IUserRepo interface {
	Register(email string, password string) int
}

// khai báo
type UserRepo struct {}


func (ur *UserRepo) Register(email string, password string) int {
	return 1
}

// khởi tạo
func NewUserRepo() IUserRepo {
	return &UserRepo{}
}
