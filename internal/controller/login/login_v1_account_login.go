package login

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "yclw/ygpay/api/login/v1"
	"yclw/ygpay/pkg/captcha"
)

func (c *ControllerV1) AccountLogin(ctx context.Context, req *v1.AccountLoginReq) (res *v1.AccountLoginRes, err error) {
	// 获取登录配置
	config, err := c.ConfigService.GetConfigByKey(ctx, "loginCaptchaSwitch")
	if err != nil {
		return
	}

	// 校验验证码
	// 当前使用默认的验证码生成及其校验方式，可以改为使用redis存储验证码
	if gconv.Bool(config) && !captcha.DefaultCaptcha.Verify(req.Cid, req.Code) {
		err = gerror.New("验证码错误")
		return
	}

	// 账号登录
	res, err = c.LoginService.AccountLogin(ctx, req.Username, req.Password)
	if err != nil {
		return
	}
	return
}
