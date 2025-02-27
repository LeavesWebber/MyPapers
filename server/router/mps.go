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
	}

	// 微信支付回调不需要认证
	Router.POST("mps/wxpay/notify", mpsApi.WxPayNotify)
} 