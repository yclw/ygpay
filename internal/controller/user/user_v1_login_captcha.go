package user

import (
	"context"
	"strconv"

	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/pkg/captcha"
)

func (c *ControllerV1) LoginCaptcha(ctx context.Context, _ *v1.LoginCaptchaReq) (res *v1.LoginCaptchaRes, err error) {
	captchaType, err := c.ConfigService.GetConfigByKey(ctx, "loginCaptchaType")
	if err != nil {
		return
	}
	captchaTypeInt, err := strconv.Atoi(captchaType)
	if err != nil {
		return
	}
	cid, base64 := captcha.Generate(ctx, captchaTypeInt)
	res = &v1.LoginCaptchaRes{Cid: cid, Base64: base64}
	return
}
