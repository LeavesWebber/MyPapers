package config

type Redis struct {
	Addr        string `mapstructure:"addr"`
	Password    string `mapstructure:"password"`
	DB          int    `mapstructure:"db"`
	ExpiredTime int    `mapstructure:"expired_time"`
}
