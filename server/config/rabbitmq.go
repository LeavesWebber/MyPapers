package config

type RabbitMQ struct {

	//队列名称
	Queuename string `mapstructure:"queue-name"`

	//交换机名称
	Exchange string `mapstructure:"exchange"`
	//bind Key 名称
	Key string `mapstructure:"key"`
	//连接信息
	Mqurl string `mapstructure:"mqurl"`
}
