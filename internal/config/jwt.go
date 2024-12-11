package config

type Jwt struct {
	Secret               string `mapstructure:"secret"`
	Ttl                  int64  `mapstructure:"ttl"`
	BlackListGracePeriod int64  `mapstructure:"black_list_grace_period"`
	RefreshGracePeriod   int64  `mapstructure:"refresh_grace_period"`
}
