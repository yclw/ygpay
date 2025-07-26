package login

import (
	"context"

	v1 "yclw/ygpay/api/login/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) LoginConfig(ctx context.Context, req *v1.LoginConfigReq) (res *v1.LoginConfigRes, err error) {
	confs, err := c.ConfigService.GetGroupConfig(ctx, "login")
	if err != nil {
		return
	}
	res = &v1.LoginConfigRes{
		RegisterSwitch: gconv.Bool(confs["loginRegisterSwitch"]),
		CaptchaSwitch:  gconv.Bool(confs["loginCaptchaSwitch"]),
		CaptchaType:    gconv.Int(confs["loginCaptchaType"]),
		Protocol:       gconv.String(confs["loginProtocol"]),
		Policy:         gconv.String(confs["loginPolicy"]),
	}
	return
}
