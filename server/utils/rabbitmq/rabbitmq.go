package rabbitmq

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"server/global"
)

// 断开channel 和 connection
func Destory() {
	global.MPS_RABBITMQ.Channel.Close()
	global.MPS_RABBITMQ.Conn.Close()
}

// 直接模式队列生产
func PublishSimple(message []byte) {

	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := global.MPS_RABBITMQ.Channel.QueueDeclare(
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
	}
	//调用channel 发送消息到队列中
	global.MPS_RABBITMQ.Channel.Publish(
		global.MPS_CONFIG.RabbitMQ.Exchange,
		global.MPS_CONFIG.RabbitMQ.Queuename,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			DeliveryMode: 2, //持久化
			ContentType:  "application/json",
			Body:         message,
		})

}

// simple 模式下消费者
func ConsumeSimple() <-chan amqp.Delivery {

	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := global.MPS_RABBITMQ.Channel.QueueDeclare(
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
	}

	//接收消息
	msgs, err := global.MPS_RABBITMQ.Channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		global.MPS_LOG.Error("rabbitmq consume fail", zap.Error(err))
	}
	return msgs

}
