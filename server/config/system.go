package config

type System struct {
	Env           string `mapstructure:"env"`            // 环境值
	Addr          int    `mapstructure:"addr"`           // 端口值
	DbType        string `mapstructure:"db-type"`        // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type"`       // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint"` // 多点登录拦截
	UseRedis      bool   `mapstructure:"use-redis"`      // 使用redis
	LimitCountIP  int    `mapstructure:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time"`
	StartTime     string `mapstructure:"start-time"`
	MachineID     int64  `mapstructure:"machine-id"`
}
