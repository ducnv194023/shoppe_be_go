package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ducnv194023/shoppe_be_go/internal/database/sqlc"
)

type IUserRepo interface {
	GetUserBasicByEmail(
		ctx context.Context,
		email string,
	) (*database.UserBasic, error)
	CreateUserBasic()
	CreateUserProfile()
}

type UserRepo struct{
	queries *database.Queries
}

func (ur *UserRepo) CreateUserBasic() {
	panic("unimplemented")
}

func (ur *UserRepo) CreateUserProfile() {
	panic("unimplemented")
}

func (ur *UserRepo) GetUserBasicByEmail(ctx context.Context, email string) (*database.UserBasic, error) {
    user, err := ur.queries.GetUserBasicByEmail(ctx, email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("email is not exist")
        }
        return nil, fmt.Errorf("failed to get user by email: %w", err)
    }
    
    return &user, nil
}

func NewUserRepo(queries *database.Queries) *UserRepo {
    return &UserRepo{
        queries: queries,
    }
}
