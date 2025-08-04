package casbin

import (
	"context"

	v1 "yclw/ygpay/api/casbin/v1"
)

func (c *ControllerV1) RemovePolicy(ctx context.Context, req *v1.RemovePolicyReq) (res *v1.RemovePolicyRes, err error) {
	err = c.CasbinService.RemovePolicy(ctx, req.RoleId, req.Path, req.Method)
	if err != nil {
		return nil, err
	}
	return &v1.RemovePolicyRes{}, nil
}
