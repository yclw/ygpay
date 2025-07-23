package sys

import (
	"context"

	v1 "yclw/ygpay/api/sys/v1"
)

func (c *ControllerV1) GetGroupConfig(ctx context.Context, req *v1.GetGroupConfigReq) (res *v1.GetGroupConfigRes, err error) {
	confs, err := c.ConfigService.GetGroupConfig(ctx, req.Group)
	if err != nil {
		return
	}
	res = &v1.GetGroupConfigRes{
		List: confs,
	}
	return
}
