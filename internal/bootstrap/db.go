package bootstrap

import (
	"gin-frame-base/internal/database/mysql"
	"gorm.io/gorm"
)

func initMysqlGorm() *gorm.DB {
	return mysql.GetConnection()
}
