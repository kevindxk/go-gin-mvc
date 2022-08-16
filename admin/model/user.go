/**
 * @Author: dexukong
 * @Description:
 * @File:  login
 * @Version: 1.0.0
 * @Date: 2022/08/08 10:19
 */

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:name;type:varchar(255);"json:"username"`
	Password string `gorm:"column:name;type:varchar(255);"json:"password"`
	Phone    string `gorm:"column:name;type:varchar(20);"json:"phone"`
	Email    string `gorm:"column:name;type:varchar(255);"json:"email""`
}

func (table User) TabkeName() string {
	return "user"
}
