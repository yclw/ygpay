package sys

import (
	"context"

	v1 "yclw/ygpay/api/sys/v1"
)

func (c *ControllerV1) GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error) {
	confs, err := c.ConfigService.GetConfig(ctx)
	if err != nil {
		return
	}
	res = &v1.GetConfigRes{
		List: confs,
	}
	return
}
