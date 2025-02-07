package bootstrap

import (
	cosint "gin-frame-base/internal/file_system/cos"
	"gin-frame-base/internal/global"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func InitFileSystem() {
	driver := global.Config.FileSystem.Driver
	switch driver {
	case "cos":
		global.Cos = initCos()
	default:
	}
}

func initCos() (c *cos.Client) {
	return cosint.GetConnection()
}
