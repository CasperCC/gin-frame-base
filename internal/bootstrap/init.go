package bootstrap

import (
	"gin-frame-base/internal/config"
	"gin-frame-base/internal/database/mysql"
	"gin-frame-base/internal/global"
)

func init() {
	// 初始化配置文件
	global.Config = config.GetConfig()

	global.Db = mysql.GetConnection()
}
