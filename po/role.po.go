package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          int   `gorm:"primarykey; autoIncrement"`
	Name        string
	Description string
}

func (r *Role) TableName() string {
	return "go_db_role"
}
