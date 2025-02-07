package global

import (
	"gin-frame-base/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Db     *gorm.DB
	Logger *zap.Logger
	Redis  *redis.Client
	Cos    *cos.Client
)
