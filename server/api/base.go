package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"server/global"
	"server/model/response"
)

type BaseApi struct {
}

var store = base64Captcha.DefaultMemStore // 默认内存存储实现

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
