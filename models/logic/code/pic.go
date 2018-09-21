package code

import (
	"github.com/mojocn/base64Captcha"
	"time"
)

func init() {
	store := base64Captcha.NewMemoryStore(10240, time.Minute*5)
	base64Captcha.SetCustomStore(store)
}

// 生成图片验证码
func GenerateCaptchaHandle() (mp map[string]string) {
	var config interface{}
	config = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      200,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	// 生成验证码ID， 和图片验证码的相关信息
	captchaId, digitCap := base64Captcha.GenerateCaptcha("", config)
	// 编码图片
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	//设置json响应
	mp = map[string]string{"captchaId": captchaId, "base64": base64Png}
	return
}

func VerifyCaptcha(id, code string) bool {
	return base64Captcha.VerifyCaptcha(id, code)
}
