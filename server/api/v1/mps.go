package v1

import (
	"encoding/xml"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/service"
	"server/utils"
)

type MPSApi struct{}

var mpsService = new(service.MPSService)

// CreateRechargeOrder 创建充值订单
// @Tags mps
// @Summary 创建充值订单
// @accept application/json
// @Produce application/json
// @Param data body request.CreateRechargeOrderReq true "创建充值订单"
// @Success 200 {object} response.Response{data=response.CreateRechargeOrderResp} "创建成功"
// @Router /mps/createOrder [post]
func (m *MPSApi) CreateRechargeOrder(c *gin.Context) {
	var req request.CreateRechargeOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserID(c)

	resp, err := mpsService.CreateRechargeOrder(userID, &req)
	if err != nil {
		response.FailWithMessage("创建订单失败: "+err.Error(), c)
		return
	}

	response.OkWithData(resp, c)
}

// GetOrderStatus 获取订单状态
// @Tags mps
// @Summary 获取订单状态
// @accept application/json
// @Produce application/json
// @Param order_no query string true "订单号"
// @Success 200 {object} response.Response{data=response.OrderStatusResp} "获取成功"
// @Router /mps/orderStatus [get]
func (m *MPSApi) GetOrderStatus(c *gin.Context) {
	var req request.GetOrderStatusReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	resp, err := mpsService.GetOrderStatus(req.OrderNo)
	if err != nil {
		response.FailWithMessage("获取订单状态失败: "+err.Error(), c)
		return
	}

	response.OkWithData(resp, c)
}

// GetMPSBalance 获取MPS代币余额
// 该方法通过HTTP请求获取用户钱包地址，并返回该地址的MPS代币余额
// 参数:
//
//	c *gin.Context - HTTP请求上下文，包含请求数据和响应方法
//
// 该方法没有返回值，但会通过c.JSON返回响应数据
func (m *MPSApi) GetMPSBalance(c *gin.Context) {
	// 解析请求体，获取钱包地址
	req := new(request.GetMPSBalanceReq)
	if err := c.ShouldBindJSON(req); err != nil {
		// 如果请求数据解析失败，则返回错误信息
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 调用服务层方法，根据钱包地址获取MPS余额
	resp, err := mpsService.GetMPSBalance(common.HexToAddress(req.WalletAddr))
	if err != nil {
		// 如果获取余额失败，则返回错误信息
		response.FailWithMessage("获取余额失败: "+err.Error(), c)
		return
	}

	// 如果获取余额成功，则返回余额信息
	response.OkWithData(resp, c)
}
func (m *MPSApi) GetMPSTransactions(c *gin.Context) {
	// 解析请求体，获取钱包地址
	req := new(request.GetMPSTransactionsReq)
	if err := c.ShouldBindJSON(req); err != nil {
		// 如果请求数据解析失败，则返回错误信息
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 调用服务层方法，根据钱包地址获取MPS余额
	resp, err := mpsService.GetMPSTransactions(req.UserId)
	if err != nil {
		// 如果获取余额失败，则返回错误信息
		response.FailWithMessage("获取余额失败: "+err.Error(), c)
		return
	}

	// 如果获取余额成功，则返回余额信息
	response.OkWithData(resp, c)
}
func (m *MPSApi) GetMPSRate(c *gin.Context) {
	response.OkWithData(global.MPS_CONFIG.Business.MPSExchangeRate, c)
}

// WxPayNotify 微信支付回调
// @Tags mps
// @Summary 微信支付回调
// @accept application/xml
// @Produce application/xml
// @Success 200 {string} string "成功"
// @Router /mps/wxpay/notify [post]
func (m *MPSApi) WxPayNotify(c *gin.Context) {
	// 读取请求体
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": "读取请求失败"})
		return
	}

	// 解析XML
	var notifyData map[string]string
	if err := xml.Unmarshal(body, &notifyData); err != nil {
		c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": "解析XML失败"})
		return
	}

	// 处理回调
	if err := mpsService.HandleWxPayNotify(notifyData); err != nil {
		c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": err.Error()})
		return
	}

	c.XML(http.StatusOK, gin.H{"return_code": "SUCCESS", "return_msg": "OK"})
}

// AliPayNotify 处理支付宝的异步通知。
// 该方法主要用于接收和处理来自支付宝的支付通知，验证支付的合法性，
// 并向支付宝发送接收通知的确认信息。
// 参数:
//
//	c *gin.Context: Gin框架的上下文对象，用于处理HTTP请求和响应。
func (m *MPSApi) AliPayNotify(c *gin.Context) {
	// 处理回调
	// 尝试处理支付宝的通知，如果发生错误，则响应支付宝服务器表示处理失败。
	global.MPS_LOG.Info("处理回调")
	if err := mpsService.HandleAliPayNotify(c); err != nil {
		c.String(http.StatusOK, "%s", "fail")
		return
	}
	//输出 success 表示消息获取成功，支付宝就会停止发送异步
	//如果输出 fail，表示消息获取失败，支付宝会重新发送消息到异步地址
	// 如果通知处理成功，则响应支付宝服务器表示处理成功。
	global.MPS_LOG.Info("回调成功")
	c.String(http.StatusOK, "%s", "success")
}

// BuyMPSWithFiat 处理用户使用法定货币购买MPS的请求
// 参数: c *gin.Context - Gin框架的上下文，用于处理HTTP请求和响应
func (m *MPSApi) BuyMPSWithFiat(c *gin.Context) {
	// 解析并验证请求体
	var req request.BuyMPSWithFiatReq
	// 绑定请求参数并进行错误处理
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户ID并校验
	userID := utils.GetUserID(c)

	// 调用服务并处理错误
	resp, err := mpsService.Pay(userID, &req)
	if err != nil {
		response.FailWithMessage("支付服务调用失败: "+err.Error(), c)
		return
	}

	// 返回成功结果
	response.OkWithData(resp, c)
}

// SellMPSToFiat 处理将MPS币兑换为法定货币的请求
// 该函数接收一个gin.Context参数，用于处理HTTP请求和响应
// 它从请求中提取兑换信息，验证用户身份，并调用服务层逻辑进行余额查询
func (m *MPSApi) SellMPSToFiat(c *gin.Context) {
	// 解析请求体，获取钱包地址
	req := new(request.SellMPSToFiatReq)
	if err := c.ShouldBindJSON(req); err != nil {
		// 如果请求数据解析失败，则返回错误信息
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取当前用户ID并校验
	userID := utils.GetUserID(c)
	// 调用服务层方法，根据钱包地址获取MPS余额
	resp, err := mpsService.Sell(userID, req)
	if err != nil {
		// 如果获取余额失败，则返回错误信息
		response.FailWithMessage("获取余额失败: "+err.Error(), c)
		return
	}
	// 如果获取余额成功，则返回余额信息
	response.OkWithData(resp, c)
}
