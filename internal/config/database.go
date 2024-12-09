package config

type Database struct {
	Driver         string
	Host           string
	Port           string
	Username       string
	Password       string
	Database       string
	Charset        string
	LogFileName    string
	LogMode        string
	ErrLogFileName string
	ErrLogMode     string
}
