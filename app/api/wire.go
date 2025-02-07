//go:build wireinject

package api

import (
	"gin-frame-base/app/dao"
	"gin-frame-base/app/service"
	"github.com/google/wire"
)

func InitializeUserApi() *UserApi {
	wire.Build(
		dao.NewUserDao,
		service.NewUserService,
		NewUserApi,
	)
	return &UserApi{}
}

func InitializeFileApi() *FileApi {
	wire.Build(
		dao.NewFileDao,
		service.NewFileService,
		NewFileApi,
	)
	return &FileApi{}
}
