package global

import (
	"gin-frame-base/internal/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Db     *gorm.DB
	Logger *zap.Logger
	Redis  *redis.Client
)
