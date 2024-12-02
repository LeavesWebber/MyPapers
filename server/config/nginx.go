package config

type Nginx struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
