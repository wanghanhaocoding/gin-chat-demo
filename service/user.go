package service

import (
	"gin-chat-demo/model"
	"gin-chat-demo/serializer"
)

// api来的入参
type UserRegisterService struct {
	UserName string `json:"user_name" form:"user_name" `
	Password string `json:"password" form:"password" `
}

// 绑定一个方法
func (service *UserRegisterService) Register() serializer.Response {
	var user model.User
	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).First(&user).Count(&count)
	if count != 0 {
		return serializer.Response{
			Status: 400,
			Msg:    "用户名已经存在了",
		}
	}
	user = model.User{
		UserName: service.UserName,
	}
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "加密出错了",
		}
	}
	model.DB.Create(&user)
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}
