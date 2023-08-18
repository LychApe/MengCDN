package model

import (
	"MengCDN/internal/tools/errmsg"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	User     string `gorm:"column:user; type:varchar(255);" json:"user"`
	Password string `gorm:"column:password; type:varchar(255);" json:"password"`
}

func (table User) TableName() string {
	return "User"
}

// CheckLogin 登录验证
func CheckLogin(username string, password string) (code int) {
	var user User

	DB.Where("user = ?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ErrorUserNotExist
	}
	if password != user.Password {
		return errmsg.ErrorUserPasswordWrong
	}
	return errmsg.SUCCESS

}
