package config

type Mysql struct {
	Path         string `mapstructure:"path"`           // 服务器地址:端口
	Port         string `mapstructure:"port"`           //:端口
	Config       string `mapstructure:"config"`         // 高级配置
	Dbname       string `mapstructure:"db-name"`        // 数据库名
	Username     string `mapstructure:"username"`       // 数据库用户名
	Password     string `mapstructure:"password"`       // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode"`       // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap"`        // 是否通过zap写入日志文件
}
