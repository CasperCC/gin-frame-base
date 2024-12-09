package config

type Database struct {
	Driver      string
	Host        string
	Port        string
	Username    string
	Password    string
	Database    string
	Charset     string
	LogFileName string `mapstructure:"log_file_name"`
	LogMode     string `mapstructure:"log_mode"`
}
