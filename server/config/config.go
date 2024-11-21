package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt"`
	Zap     Zap     `mapstructure:"zap"`
	System  System  `mapstructure:"system"`
	Captcha Captcha `mapstructure:"captcha"`
	Mysql   Mysql   `mapstructure:"mysql"`
	// 跨域配置
	Cors          CORS   `mapstructure:"cors" json:"cors" yaml:"cors"`
	ImagePath     string `mapstructure:"image-path"`
	ImageIpfsPath string `mapstructure:"image-ipfs-path"`
	IPFS          IPFS   `mapstructure:"ipfs"`
	Nginx         Nginx  `mapstructure:"nginx"`
}
