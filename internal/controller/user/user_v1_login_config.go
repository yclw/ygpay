package user

import (
	"context"

	v1 "yclw/ygpay/api/user/v1"
)

func (c *ControllerV1) LoginConfig(ctx context.Context, req *v1.LoginConfigReq) (res *v1.LoginConfigRes, err error) {
	confs, err := c.ConfigService.GetGroupConfig(ctx, "login")
	if err != nil {
		return
	}
	res = &v1.LoginConfigRes{
		RegisterSwitch: confs["loginRegisterSwitch"].(int),
		CaptchaSwitch:  confs["loginCaptchaSwitch"].(int),
		CaptchaType:    confs["loginCaptchaType"].(int),
		Avatar:         confs["loginAvatar"].(string),
		RoleId:         confs["loginRoleId"].(int64),
		Protocol:       confs["loginProtocol"].(string),
		Policy:         confs["loginPolicy"].(string),
	}
	return
}
