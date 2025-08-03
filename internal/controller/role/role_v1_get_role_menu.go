package role

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"yclw/ygpay/api/role/v1"
)

func (c *ControllerV1) GetRoleMenu(ctx context.Context, req *v1.GetRoleMenuReq) (res *v1.GetRoleMenuRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
