package core

import (
	"server/global"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// 创建简单模式下RabbitMQ实例
func NewRabbitMQSimple() *global.RabbitMQ {
	//获取connection
	Conn, err := amqp.Dial(global.MPS_CONFIG.RabbitMQ.Mqurl)
	if err != nil {
		global.MPS_LOG.Error("failed to connect rabbitmq", zap.Error(err))
		return nil
	}

	channel, err := Conn.Channel()
	if err != nil {
		global.MPS_LOG.Error("failed to open a channel", zap.Error(err))
		Conn.Close()
		return nil
	}

	_, err = channel.QueueDeclare(
		global.MPS_CONFIG.RabbitMQ.Queuename,
		//是否持久化
		true,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		global.MPS_LOG.Error("failed to declare a queue", zap.Error(err))
		channel.Close()
		Conn.Close()
		return nil
	}

	return &global.RabbitMQ{
		Conn:    Conn,
		Channel: channel,
	}
}
