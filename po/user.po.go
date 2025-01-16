package po

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type User struct {
	gorm.Model
	UUID      uuid.UUID `gorm:"unique; type:varchar(255); not null; index: idx_uuid;"`
	Username  string
	Password  string
	IsActive  bool
	Roles     []Role `gorm:"many2many:go_db_user_roles;"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
