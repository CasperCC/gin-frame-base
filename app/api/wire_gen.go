// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"gin-frame-base/app/dao"
	"gin-frame-base/app/service"
)

// Injectors from wire.go:

func InitializeUserApi() *UserApi {
	userDao := dao.NewUserDao()
	userService := service.NewUserService(userDao)
	userApi := NewUserApi(userService)
	return userApi
}

func InitializeFileApi() *FileApi {
	fileDao := dao.NewFileDao()
	fileService := service.NewFileService(fileDao)
	fileApi := NewFileApi(fileService)
	return fileApi
}
