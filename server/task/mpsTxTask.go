package task

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/service"
	"server/utils/rabbitmq"
)

func Task() {
	msgs := rabbitmq.ConsumeSimple()
	go func() {
		for msg := range msgs {
			var queueMsg global.QueueMessage
			err := json.Unmarshal(msg.Body, &queueMsg)
			if err != nil {
				global.MPS_LOG.Error("反序列化JSON失败", zap.Error(err))
				if nackErr := msg.Nack(false, true); nackErr != nil {
					global.MPS_LOG.Error("发送NACK以重新入队消息失败", zap.Error(nackErr))
				}
				return
			}
			fmt.Println("handle queueMsg")
			err = service.TransMpsToWallet(queueMsg.UUID, queueMsg.Address, queueMsg.MPSAmount, queueMsg.OrderNo, queueMsg.Description)
			if err != nil {
				global.MPS_LOG.Error("TransMpsToWallet失败", zap.Error(err))
				// 可根据业务需求实现重试或补偿逻辑
				if nackErr := msg.Nack(false, true); nackErr != nil {
					global.MPS_LOG.Error("发送NACK以重新入队消息失败", zap.Error(nackErr))
				}
				return
			}
			// 如果处理成功，发送 ACK
			if ackErr := msg.Ack(false); ackErr != nil {
				global.MPS_LOG.Error("发送ACK失败", zap.Error(ackErr))
			}
		}
	}()
}
