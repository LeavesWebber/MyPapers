package router

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type MPSRouter struct{}

// InitMPSRouter 初始化 MPS 路由
func (r *MPSRouter) InitMPSRouter(Router *gin.RouterGroup) {
	mpsRouter := Router.Group("mps").Use(middleware.JWTAuth())
	mpsApi := v1.ApiGroupApp.MPSApi
	{
		mpsRouter.POST("createOrder", mpsApi.CreateRechargeOrder) // 创建充值订单
		mpsRouter.GET("orderStatus", mpsApi.GetOrderStatus)       // 获取订单状态
		mpsRouter.GET("balance", mpsApi.GetMPSBalance)            // 查询用户MPS余额
		mpsRouter.GET("transactions", mpsApi.GetMPSTransactions)  // 获取用户通证交易记录
		mpsRouter.POST("buy", mpsApi.GetOrderStatus)              // 法币购买MPS
		mpsRouter.POST("sell", mpsApi.GetOrderStatus)             // MPS卖出换取法币
		mpsRouter.GET("rate", mpsApi.GetOrderStatus)              // 获取当前MPS兑换率
		mpsRouter.GET("rewards", mpsApi.GetOrderStatus)           // 获取用户激励记录
	}

	// 微信支付回调不需要认证
	Router.POST("mps/wxpay/notify", mpsApi.WxPayNotify)
	Router.POST("mps/alipay/notify", mpsApi.AliPayNotify)
}
