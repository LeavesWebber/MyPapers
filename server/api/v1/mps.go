package v1

import (
	"encoding/xml"
	"io/ioutil"
	"server/model/request"
	"server/model/response"
	"server/service"
	"server/utils"
	"net/http"

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