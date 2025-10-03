package request

type SendEmailCode struct {
	Email     string `json:"email" binding:"required,email"`
	Captcha   string `json:"captcha" binding:"required,len=6"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}
