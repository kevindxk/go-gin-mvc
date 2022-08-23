/**
 * @Author: dexukong
 * @Description:
 * @File:  userService
 * @Version: 1.0.0
 * @Date: 2022/08/23 10:09
 */

package services

import (
	"context"
	"fmt"
	"goproject/admin/model"
	"goproject/admin/pkg/e"
	"goproject/admin/serializer"
)

type UserService struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (userser *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	fmt.Println(userser)
	fmt.Println(user)
	fmt.Println(ctx)
	code := e.Success
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   userser,
	}
}

func (userser *UserService) Login(ctx context.Context) serializer.Response {
	var user model.User
	fmt.Println(userser)
	fmt.Println(user)
	fmt.Println(ctx)
	code := e.Success
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   userser,
	}
}
