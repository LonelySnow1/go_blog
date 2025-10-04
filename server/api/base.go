package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
)

/*
 	API层只做参数传递和错误处理 —— 对应 Controller/Handler
	直接接收客户端的 HTTP 请求（如 GET/POST 等），解析请求参数（路径参数、表单、JSON 等）。
	负责参数合法性校验（如必填项、格式校验），过滤无效或恶意输入。
	调用 Service 层执行具体业务逻辑，接收返回结果。
	将 Service 层的处理结果封装为标准响应（如 JSON 格式），返回给客户端（包括正常数据或错误信息）。
*/

type BaseApi struct {
}

var store = base64Captcha.DefaultMemStore // 默认内存存储实现

// Captcha 创建图形验证码
func (baseApi *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit( // 创建数字验证码的驱动
		global.Config.Captcha.Height,   // 高
		global.Config.Captcha.Width,    // 宽
		global.Config.Captcha.Length,   // 长
		global.Config.Captcha.MaxSkew,  // 倾斜程度/扭曲程度
		global.Config.Captcha.DotCount, // 干扰点数量
	)

	captcha := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := captcha.Generate()

	if err != nil {
		global.Log.Error("Failed to generate captcha", zap.Error(err))
		response.FailWithMessage("Failed to generate captcha", c)
		return
	}
	response.OkWithData(response.Captcha{
		CaptchaId: id,
		PicPath:   b64s,
	}, c)
}

// SendEmailCode 发送邮箱验证码
func (baseApi *BaseApi) SendEmailCode(c *gin.Context) {
	var req request.SendEmailCode
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(req.CaptchaId, req.Captcha, true) { // 判断验证码是否正确 clear : 是否验证一次就清除 true - 是

		err := baseService.SendEmailCode(c, req.Email)

		if err != nil {
			global.Log.Error("Failed to send email", zap.Error(err))
			response.FailWithMessage("Failed to send email", c)
			return
		}

		response.OkWithMessage("successfully send email", c)
		return
	}
	response.FailWithMessage("Incorrect verification code", c)
}

// QQLoginURL 返回 QQ 登录链接
func (baseApi *BaseApi) QQLoginURL(c *gin.Context) {
	url := global.Config.QQ.QQLoginURL()
	response.OkWithData(url, c)
}
