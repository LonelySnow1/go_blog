package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/utils"
	"time"
)

type BaseService struct {
}

func (baseService *BaseService) SendEmailCode(c *gin.Context, to string) error {
	verificationCode := utils.GenerateVerificationCode(6)
	expireTime := time.Now().Add(5 * time.Minute).Unix() // 过期时间

	session := sessions.Default(c)
	session.Set("verification_code", verificationCode)
	session.Set("expire_time", expireTime)
	session.Set("email", to)
	_ = session.Save()

	subject := "喵~你的身份验证码来啦！"
	body := `<body style="font-family: 'Arial Rounded MT Bold', 'Helvetica Neue', sans-serif; line-height: 1.6; color: #555;">
    <p>亲爱的<span style="color: #ff6b8b; font-weight: bold;">` + to + `酱</span>：</p>

    <p>今天有没有赖床呀？嘻嘻，感谢主人注册` + global.Config.Website.Name + `的个人博客呢，本喵来给你送验证码了喵！</p>

    <p>您这次的专属验证码是：<br>
    <span style="color: #ff7a00; font-size: 1.2em; font-weight: bold; letter-spacing: 2px;">【` + verificationCode + `】</span></p>

    <p>这个验证码只有5分钟的有效期哦！过了时间就会“失效”啦，所以一定要抓紧时间用它完成验证呀～</p>

    <p>另外要偷偷提醒你：这个验证码只能自己看、自己用，绝对不能告诉其他陌生人喵！要是不小心弄丢了，也可以重新申请哦～</p>

    <p>希望<span style="color: #ff6b8b; font-weight: bold;">` + to + `酱</span>能顺顺利利完成验证，要是有任何问题，都可以找小助手帮忙哦～</p>

    <p>啾咪～❤️<br>
    <span style="color: #8a5cf7;">猫娘小助手の邮箱:` + global.Config.Email.From + `</span></p>

    <p style="font-size: 0.9em; color: #999;">—— ` + global.Config.Website.Title + `&nbsp;今天也有好好营业的小猫咪呀～</p>
</body>`

	_ = utils.Email(to, subject, body)

	return nil
}
