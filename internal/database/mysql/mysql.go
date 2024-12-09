package mysql

import (
	"gin-frame-base/internal/config"
	"gin-frame-base/internal/global"
	"gin-frame-base/internal/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// GetConnection 连接mysql
func GetConnection() (db *gorm.DB) {
	// 获取数据库配置
	database := global.Config.Database

	// 拼装dsn
	dsn := database.Username + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.Database + "?charset=" + database.Charset + "&parseTime=True&loc=Local"

	// 初始化配置
	mysqlConfig := mysql.Config{
		DSN: dsn,
	}

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 禁用表名复数
		},
		Logger: getGormLogger(),
	})
	if err != nil {
		log.Println("数据库链接失败：", err)
		panic(err)
	}

	return
}

func getGormLogger() gormLogger.Interface {
	gormLoggerConfig := config.Logger{
		Level:      global.Config.Database.LogMode,
		FilePath:   global.Config.Logger.FilePath,
		FileName:   global.Config.Database.LogFileName,
		MaxSize:    global.Config.Logger.MaxSize,
		MaxAge:     global.Config.Logger.MaxAge,
		MaxBackups: global.Config.Logger.MaxBackups,
		Compress:   true,
	}
	gLogger, err := logger.New(gormLoggerConfig)
	if err != nil {
		log.Println("数据库日志初始化失败：", err)
		panic(err)
	}
	return gormLogger.New(
		zap.NewStdLog(gLogger),
		gormLogger.Config{
			SlowThreshold: time.Second,                                 // 慢查询阈值
			LogLevel:      ConvertZapLevelToGormLevel(gLogger.Level()), // Gorm 日志级别
			Colorful:      false,                                       // 禁用彩色日志
		},
	)
}

func ConvertZapLevelToGormLevel(zapLevel zapcore.Level) gormLogger.LogLevel {
	switch zapLevel {
	case zapcore.DebugLevel:
		return gormLogger.Info // Gorm 没有 Debug 级别，映射为 Info
	case zapcore.InfoLevel:
		return gormLogger.Info
	case zapcore.WarnLevel:
		return gormLogger.Warn
	case zapcore.ErrorLevel:
		return gormLogger.Error
	default:
		return gormLogger.Info // 默认使用 Gorm 的 Info 级别
	}
}
