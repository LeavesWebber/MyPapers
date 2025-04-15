package task

import (
	"encoding/json"
	"fmt"
	"server/global"
	"server/model/tables"
	"server/service"
	Alipay "server/utils/alipay"
	"server/utils/rabbitmq"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func Task() {
	MPSTxTask()
	MPSRechargeOrderTask()
}

// table MPS_Transition
func MPSTxTask() {
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
			global.MPS_LOG.Info("handle queueMsg" + queueMsg.Description)
			err = service.TransMpsToWallet(queueMsg.UUID, queueMsg.Address, queueMsg.MPSAmount, queueMsg.OrderNo, queueMsg.Description)
			if err != nil {
				global.MPS_LOG.Error("TransMpsToWallet失败", zap.Error(err))
				// 可根据业务需求实现重试或补偿逻辑
				if nackErr := msg.Nack(false, false); nackErr != nil {
					global.MPS_LOG.Error("发送NACK消息失败", zap.Error(nackErr))
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

// table MPS_Recharge_Order
func MPSRechargeOrderTask() {
	// 在开发环境中不执行订单清理任务
	if global.MPS_CONFIG.System.Env == "develop" {
		global.MPS_LOG.Info("开发环境，跳过订单清理任务")
		return
	}

	//如果还在执行则跳过
	job := cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
	job.AddFunc("@every 5m", func() {
		var tx []tables.MPSRechargeOrder
		err := global.MPS_DB.Where("status = ?", 0).Find(&tx).Limit(20).Error
		if err != nil {
			global.MPS_LOG.Error("查询MPS_Recharge_Order失败", zap.Error(err))
			return
		}

		//如果超时则删除
		for _, v := range tx {
			timeOut := time.Duration(global.MPS_CONFIG.Business.OrderTimeout) * time.Minute
			if v.CreatedAt.Add(timeOut).Before(time.Now()) {
				// 尝试关闭支付宝订单
				err := Alipay.TradeClose("nil", v.OrderNo)
				if err != nil {
					// 如果是系统错误，记录警告日志
					if err.Error() == "支付宝系统错误" {
						global.MPS_LOG.Warn(fmt.Sprintf("支付宝关闭订单号: %v失败，系统繁忙", v.OrderNo))
					} else {
						global.MPS_LOG.Error(fmt.Sprintf("支付宝关闭订单号: %v失败", v.OrderNo), zap.Error(err))
					}
					continue
				}

				// 删除订单记录
				err = global.MPS_DB.Where("id = ?", v.ID).Delete(&tables.MPSRechargeOrder{}).Error
				if err != nil {
					global.MPS_LOG.Error(fmt.Sprintf("MPS_Recharge_Order 删除订单号: %v失败", v.OrderNo), zap.Error(err))
					continue
				}
				global.MPS_LOG.Info(fmt.Sprintf("MPS_Recharge_Order 删除订单号成功: %v", v.OrderNo))
				//todo 通知用户
			}
		}
	})
	job.Start()
}
