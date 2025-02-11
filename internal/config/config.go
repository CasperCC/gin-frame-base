package config

import (
	"fmt"
	"gin-frame-base/internal/tool"
	"github.com/spf13/viper"
)

type Config struct {
	App        App
	Database   Database
	Redis      Redis
	Logger     Logger
	Jwt        Jwt
	FileSystem FileSystem `mapstructure:"file_system"`
}

// GetConfig 读取配置文件
func GetConfig() (c *Config) {
	// 获取项目的根目录
	rootDir, _ := tool.GetRootDir()

	// 实例化viper，并根据地址读取配置文件
	v := viper.New()
	v.SetConfigFile(rootDir + "/config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("ReadInConfigError: ", err)
		return
	}

	// 将读取到的配置文件绑定到返回参数c
	err = v.Unmarshal(&c)
	if err != nil {
		fmt.Println("ConfigUnmarshalError: ", err)
		return
	}

	return
}
