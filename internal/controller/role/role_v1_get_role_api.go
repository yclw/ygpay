package role

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"yclw/ygpay/api/role/v1"
)

func (c *ControllerV1) GetRoleApi(ctx context.Context, req *v1.GetRoleApiReq) (res *v1.GetRoleApiRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
