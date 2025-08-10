package casbin

import (
	"context"

	v1 "yclw/ygpay/api/casbin/v1"
)

func (c *ControllerV1) RefreshEnforcer(ctx context.Context, req *v1.RefreshEnforcerReq) (res *v1.RefreshEnforcerRes, err error) {
	err = c.CasbinService.RefreshEnforcer(ctx)
	return
}
