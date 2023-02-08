package user

import (
	"goblog/app/models"
	"goblog/pkg/password"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name            string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password        string `gorm:"type:varchar(255)" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

// ComparePassword 方法用来验证用户密码 hash 值
func (u *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, u.Password)
}

// Link 方法用来生成用户链接
func (u *User) Link() string {
	return ""
}
