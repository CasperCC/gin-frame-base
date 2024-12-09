package service

import (
	_ "gin-frame-base/internal/bootstrap"
	"gin-frame-base/internal/global"
)

var (
	BaseService = &baseService{db: global.Db}

	UserService = &userService{baseService: *BaseService}
)
