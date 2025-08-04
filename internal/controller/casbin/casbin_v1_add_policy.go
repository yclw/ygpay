package casbin

import (
	"context"

	v1 "yclw/ygpay/api/casbin/v1"
)

func (c *ControllerV1) AddPolicy(ctx context.Context, req *v1.AddPolicyReq) (res *v1.AddPolicyRes, err error) {
	err = c.CasbinService.AddPolicy(ctx, req.RoleId, req.Path, req.Method)
	if err != nil {
		return nil, err
	}
	return &v1.AddPolicyRes{}, nil
}
