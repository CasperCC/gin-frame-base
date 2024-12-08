package global

import (
	"gin-frame-base/internal/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Db     *gorm.DB
)
