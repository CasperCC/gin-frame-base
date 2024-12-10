package config

type Jwt struct {
	Secret string `mapstructure:"secret"`
	Ttl    int64  `mapstructure:"ttl"`
}
