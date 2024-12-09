package config

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	Prefix   string `mapstructure:"prefix"`
}
