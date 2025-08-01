package role

import (
	"context"

	v2 "yclw/ygpay/api/role/v2"
)

func (c *ControllerV2) GetOne(ctx context.Context, req *v2.GetOneReq) (res *v2.GetOneRes, err error) {
	model, err := c.RoleService.GetOne(ctx, req.Id)
	if err != nil {
		return
	}
	res = &v2.GetOneRes{
		RoleModel: c.roleModelToV2(model),
	}
	return
}
