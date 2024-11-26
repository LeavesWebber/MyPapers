package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"server/global"
	"server/model/response"
)

var store = base64Captcha.DefaultMemStore

type UserApi struct{}

// Captcha 生成验证码
func (u *UserApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.MPS_CONFIG.Captcha.ImgHeight, global.MPS_CONFIG.Captcha.ImgWidth, global.MPS_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.MPS_LOG.Error("cp.Generate() failed", zap.Error(err))
		ResponseError(c, CodeCaptchaFailed)
	} else {
		ResponseSuccess(c, response.SysCaptcha{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.MPS_CONFIG.Captcha.KeyLong,
		})
	}
}
