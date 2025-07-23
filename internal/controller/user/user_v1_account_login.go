package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/util/captcha"
)

func (c *ControllerV1) AccountLogin(ctx context.Context, req *v1.AccountLoginReq) (res *v1.AccountLoginRes, err error) {
	// 获取登录配置
	config, err := c.ConfigService.GetConfigByKey(ctx, "loginCaptchaSwitch")
	if err != nil {
		return
	}

	// 校验验证码
	if !req.IsLock && config == "1" && !captcha.Verify(req.Cid, req.Code) {
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
