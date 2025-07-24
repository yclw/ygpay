package user

import (
	"context"

	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/pkg/captcha"
)

func (c *ControllerV1) LoginCaptcha(ctx context.Context, _ *v1.LoginCaptchaReq) (res *v1.LoginCaptchaRes, err error) {
	// 当前使用默认的验证码生成及其校验方式，可以改为使用redis存储验证码
	cid, base64, _, err := captcha.DefaultCaptcha.Generate(ctx)
	if err != nil {
		return
	}
	res = &v1.LoginCaptchaRes{Cid: cid, Base64: base64}
	return
}
