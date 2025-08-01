package role

import (
	"context"

	v2 "yclw/ygpay/api/role/v2"
)

func (c *ControllerV2) Delete(ctx context.Context, req *v2.DeleteReq) (res *v2.DeleteRes, err error) {
	err = c.RoleService.Delete(ctx, req.Id)
	if err != nil {
		return
	}
	res = &v2.DeleteRes{}
	return
}
