package config

type IPFS struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	GatewayPort string `mapstructure:"gateway-port"`
	GatewayPath string `mapstructure:"gateway-path"`
}
