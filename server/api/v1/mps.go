package v1

import (
	"encoding/xml"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"net/http"
	"server/model/request"
	"server/model/response"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type MPSApi struct{}

var mpsService = new(service.MPSService)

// CreateRechargeOrder 创建充值订单
// @Tags MPS
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
// @Tags MPS
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

// WxPayNotify 微信支付回调
// @Tags MPS
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
	if err := mpsService.HandleAliPayNotify(c); err != nil {
		c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": err.Error()})
		return
	}
	//输出 success 表示消息获取成功，支付宝就会停止发送异步
	//如果输出 fail，表示消息获取失败，支付宝会重新发送消息到异步地址
	// 如果通知处理成功，则响应支付宝服务器表示处理成功。
	c.XML(http.StatusOK, gin.H{"return_code": "SUCCESS", "return_msg": "OK"})
}
