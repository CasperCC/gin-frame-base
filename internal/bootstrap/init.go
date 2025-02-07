package bootstrap

import (
	"gin-frame-base/internal/cache/redis"
	"gin-frame-base/internal/config"
	"gin-frame-base/internal/global"
	"gin-frame-base/internal/logger"
)

func init() {
	// 初始化配置文件
	global.Config = config.GetConfig()

	var err error
	global.Logger, err = logger.New(global.Config.Logger)
	if err != nil {
		panic("初始化日志错误：" + err.Error())
		return
	}

	global.Db = InitDBGorm()
	global.Redis = redis.InitializeRedis()
	InitFileSystem()
	global.Logger.Info("server start!")
}
