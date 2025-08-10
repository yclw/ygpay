package casbin

import (
	"context"

	v1 "yclw/ygpay/api/casbin/v1"
)

func (c *ControllerV1) SyncRoleApi(ctx context.Context, req *v1.SyncRoleApiReq) (res *v1.SyncRoleApiRes, err error) {
	err = c.CasbinService.SyncRoleApi(ctx)
	return
}
