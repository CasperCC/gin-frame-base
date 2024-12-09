package bootstrap

import (
	"gin-frame-base/internal/database/mysql"
	"gin-frame-base/internal/global"
	"gorm.io/gorm"
)

func InitDBGorm() *gorm.DB {
	driver := global.Config.Database.Driver
	switch driver {
	case "mysql":
		return initMysqlGorm()
	default:
		return initMysqlGorm()
	}
}

func initMysqlGorm() *gorm.DB {
	return mysql.GetConnection()
}
