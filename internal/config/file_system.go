package config

type FileSystem struct {
	Driver     string `mapstructure:"driver"`
	SecretId   string `mapstructure:"secret_id"`
	SecretKey  string `mapstructure:"secret_key"`
	Region     string `mapstructure:"region"`
	Endpoint   string `mapstructure:"endpoint"`
	BucketName string `mapstructure:"bucket_name"`
	Prefix     string `mapstructure:"prefix"`
}
