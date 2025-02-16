package captcha_api

import (
	"fast_gin/utils/captcha"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type GenerateResponse struct {
	CaptchaId string `json:"captchaID"`
	Captcha   string `json:"captcha"`
}

func (CaptchaApi) GenerateView(c *gin.Context) {
	var driver = base64Captcha.DriverString{
		Width:           200,
		Height:          40,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          4,
		Source:          "0123456789",
	}
	cp := base64Captcha.NewCaptcha(&driver, captcha.CaptchaStore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		logrus.Errorf("验证码生成失败 %s", err)
		res.FailWithMsg("验证码生成失败", c)
		return
	}
	res.OkWithData(GenerateResponse{
		CaptchaId: id,
		Captcha:   b64s,
	}, c)
}
