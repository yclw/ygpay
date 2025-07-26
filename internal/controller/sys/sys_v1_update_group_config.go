package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "yclw/ygpay/api/sys/v1"
)

func (c *ControllerV1) UpdateGroupConfig(ctx context.Context, req *v1.UpdateGroupConfigReq) (res *v1.UpdateGroupConfigRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
