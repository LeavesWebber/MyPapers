package config

type Captcha struct {
	KeyLong   int `mapstructure:"key-long"`   // 验证码长度
	ImgWidth  int `mapstructure:"img-width"`  // 验证码宽度
	ImgHeight int `mapstructure:"img-height"` // 验证码高度
}
