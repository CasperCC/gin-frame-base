package global

import (
	"gin-frame-base/internal/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Db     *gorm.DB
	Logger *zap.Logger
)
